package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/juliotorresmoreno/searchbar/db"
	"github.com/juliotorresmoreno/searchbar/models"
)

//GetDescribe Metodo Get encargado de describir e incrementar el score de un elemento
func GetDescribe(w http.ResponseWriter, r *http.Request) {
	var data []byte
	id := r.URL.Query().Get("id")
	cache := db.GetCache()
	result := cache.Get(id)
	row := models.Datatable{}
	val := result.Val()
	if val != "" {
		response := newResponseDescribe()
		json.Unmarshal([]byte(val), &row)
		response.Data = row
		data, _ = json.Marshal(response)
	} else {
		response := newResponseDescribeNotFound()
		json.Unmarshal([]byte(val), &row)
		data, _ = json.Marshal(response)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

type responseDescribe struct {
	Success bool
	Data    models.Datatable
}

func newResponseDescribe() responseDescribe {
	return responseDescribe{
		Success: true,
	}
}

func newResponseDescribeNotFound() responseDescribeNotFound {
	return responseDescribeNotFound{
		Success: true,
		Error:   "No se encuentra el elemento",
	}
}

type responseDescribeNotFound struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
