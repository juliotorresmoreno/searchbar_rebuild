package api

import (
	"encoding/json"

	"strings"

	"github.com/juliotorresmoreno/searchbar/db"
	"github.com/juliotorresmoreno/searchbar/models"
	redis "gopkg.in/redis.v5"
)

//DescribeElement describe el elemento
func DescribeElement(id, stores string) (models.Datatable, error) {
	return models.Datatable{}, nil
}

/*
//DescribeElement describe el elemento
func DescribeElement(id, stores string) (models.Datatable, error) {
	return models.Datatable{}, nil
	_stores := strings.Split(stores, ",")
	cache := db.GetCache()
	defer cache.Close()
	for _, v := range _stores {
		if v != "" {
			incrementScore(cache, id, v)
		}
	}
	result := cache.Get(id)
	row := models.Datatable{}
	val := result.Val()
	return row, nil
	if val != "" {
		json.Unmarshal([]byte(val), &row)
		return row, nil
	}
	return row, fmt.Errorf("Not found")
}
*/

//LocationQuery Metodo Get encargado de consultar los datos
func LocationQuery(query string) ResponseLocation {
	cache := db.GetCache()
	keys := cache.FindKeys(strings.ToLower(query))
	_result := make([]z, 0)
	stores := map[string][]string{}
	proccess := map[string]bool{}
	scores := map[string]float64{}
	for _, v := range keys {
		data := cache.ZRangeWithScores(v, 0, 10)
		result, _ := data.Result()
		for i := range result {
			member := result[i].Member.(string)
			tmp := z{Id: v, Z: result[i]}
			stores[member] = append(stores[member], v)
			if v, ok := scores[member]; !ok || v < tmp.Score {
				scores[member] = tmp.Score
			}
			_result = append(_result, tmp)
		}
	}
	length := len(_result)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if _result[i].Score < _result[j].Score {
				tmp := _result[i]
				_result[i] = _result[j]
				_result[j] = tmp
			}
		}
	}
	response := newResponseLocation()
	for i := 0; /*i < 5 && */ i < length; i++ {
		member := _result[i].Member.(string)
		if _, ok := proccess[member]; !ok {
			tmp := cache.Get(member)
			row := models.Datatable{}
			json.Unmarshal([]byte(tmp.Val()), &row)
			el := newResponseLocationItem()
			el.Id = member
			el.Score = scores[member]
			el.Stores = stores[member]
			el.Datatable = row
			response.Data = append(response.Data, el)
			proccess[member] = true
		}
	}
	return response
}

type z struct {
	Id string
	redis.Z
}

//ResponseLocation Respuesta de la consulta
type ResponseLocation struct {
	Success bool                   `json:"success"`
	Data    []ResponseLocationItem `json:"data"`
}

//ResponseLocationItem Elemento
type ResponseLocationItem struct {
	Id               string   `json:"id"`
	Score            float64  `json:"score"`
	Stores           []string `json:"stores"`
	models.Datatable `json:"data"`
}

func newResponseLocation() ResponseLocation {
	return ResponseLocation{
		Success: true,
		Data:    make([]ResponseLocationItem, 0),
	}
}
func newResponseLocationItem() ResponseLocationItem {
	return ResponseLocationItem{
		Stores: make([]string, 0),
	}
}

func incrementScore(cache *db.Cache, id, store string) {
	data := cache.ZRangeWithScores(store, 0, -1)
	words, _ := data.Result()
	for _, t := range words {
		if t.Member == id {
			cache.ZAdd(store, redis.Z{
				Score:  t.Score + 1,
				Member: id,
			})
		}
	}
}
