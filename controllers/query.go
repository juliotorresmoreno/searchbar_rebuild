package controllers

import (
	"net/http"

	redis "gopkg.in/redis.v5"

	"encoding/json"

	"github.com/juliotorresmoreno/searchbar/db"
	"github.com/juliotorresmoreno/searchbar/models"
)

//GetQuery Metodo Get encargado de consultar los datos
func GetQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	cache := db.GetCache()
	keys := cache.FindKeys(query)
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
	response := newResponseQuery()
	for i := 0; i < 5 && i < length; i++ {
		member := _result[i].Member.(string)
		if _, ok := proccess[member]; !ok {
			tmp := cache.Get(member)
			row := models.Datatable{}
			json.Unmarshal([]byte(tmp.Val()), &row)
			el := newResponseQueryItem()
			el.Id = member
			el.Score = scores[member]
			el.Stores = stores[member]
			el.Datatable = row
			response.Data = append(response.Data, el)
			proccess[member] = true
		}
	}
	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

type z struct {
	Id string
	redis.Z
}

type responseQuery struct {
	Success bool
	Data    []responseQueryItem
}

type responseQueryItem struct {
	Id     string   `json:"id"`
	Score  float64  `json:"score"`
	Stores []string `json:"stores"`
	models.Datatable
}

func newResponseQuery() responseQuery {
	return responseQuery{
		Success: true,
		Data:    make([]responseQueryItem, 0),
	}
}
func newResponseQueryItem() responseQueryItem {
	return responseQueryItem{
		Stores: make([]string, 0),
	}
}
