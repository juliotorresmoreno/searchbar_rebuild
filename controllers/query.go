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
	for _, v := range keys {
		data := cache.ZRangeWithScores(v, 0, 10)
		result, _ := data.Result()
		for i := range result {
			tmp := z{Id: v, Z: result[i]}
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
		tmp := cache.Get(member)
		row := models.Datatable{}
		json.Unmarshal([]byte(tmp.Val()), &row)
		el := responseQueryItem{
			Id:        member,
			Score:     _result[i].Score,
			Store:     _result[i].Id,
			Datatable: row,
		}
		response.Data = append(response.Data, el)
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
	Id    string  `json:"id"`
	Score float64 `json:"score"`
	Store string  `json:"store"`
	models.Datatable
}

func newResponseQuery() responseQuery {
	return responseQuery{
		Success: true,
		Data:    make([]responseQueryItem, 0),
	}
}
