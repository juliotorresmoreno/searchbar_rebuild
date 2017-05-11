package controllers

import (
	"net/http"

	"encoding/json"

	"github.com/juliotorresmoreno/searchbar/api"
)

//GetQuery Metodo Get encargado de consultar los datos
func GetQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	response := api.LocationQuery(query)
	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
