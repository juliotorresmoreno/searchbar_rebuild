package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/juliotorresmoreno/searchbar/api"
	"github.com/juliotorresmoreno/searchbar/models"
)

//GetDescribe Metodo Get encargado de describir e incrementar el score de un elemento
func GetDescribe(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	stores := r.URL.Query().Get("stores")
	data, err := api.DescribeElement(id, stores)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusOK)
		response := newResponseDescribeNotFound()
		_response, _ := json.Marshal(response)
		w.Write(_response)
	} else {
		w.WriteHeader(http.StatusOK)
		response := newResponseDescribe()
		response.Data = data
		_response, _ := json.Marshal(response)
		w.Write(_response)
	}
}

type responseDescribe struct {
	Success bool             `json:"success"`
	Data    models.Datatable `json:"data"`
}

func newResponseDescribe() responseDescribe {
	return responseDescribe{
		Success: true,
	}
}

func newResponseDescribeNotFound() responseDescribeNotFound {
	return responseDescribeNotFound{
		Success: true,
		Error:   "Not found",
	}
}

type responseDescribeNotFound struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
