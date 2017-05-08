package controllers

import (
	"net/http"

	redis "gopkg.in/redis.v5"

	"encoding/json"

	"github.com/juliotorresmoreno/searchbar/db"
)

//GetQuery Metodo Get encargado de consultar los datos
func GetQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	cache := db.GetCache()
	keys := cache.FindKeys(query)
	_result := make([]redis.Z, 0)
	for _, v := range keys {
		data := cache.ZRangeWithScores(v, 0, 10)
		result, _ := data.Result()
		for i := range result {
			_result = append(_result, result[i])
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
	response := make([]map[string]string, 0)
	for i := 0; i < 5 && i < length; i++ {
		member := _result[i].Member.(string)
		tmp := cache.Get(member).Val()
		row := map[string]string{}
		json.Unmarshal([]byte(tmp), &row)
		response = append(response, row)
		println(member, tmp, cache.Get(member).Err().Error())
	}
	data, _ := json.Marshal(keys)
	w.Write(data)
}
