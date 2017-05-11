package controllers

import (
	"net/http"

	"encoding/json"

	"fmt"

	"github.com/juliotorresmoreno/searchbar/config"
	"github.com/juliotorresmoreno/searchbar/etl"
)

//GetEtl Metodo Get encargado de importar los datos
func GetEtl(w http.ResponseWriter, r *http.Request) {
	tool := etl.NewEtl()
	tool.Open()
	defer tool.Close()
	errs := make([]string, 0)
	for _, v := range config.SOURCES {
		err := tool.ImportData(v.Host, v.Database, v.Collection)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if len(errs) == 0 {
		response := responseEtl{Success: true}
		data, _ := json.Marshal(response)
		fmt.Fprintf(w, string(data))
	} else {
		response := responseEtlError{
			Success: true,
			Errors:  errs,
		}
		data, _ := json.Marshal(response)
		fmt.Fprintf(w, string(data))
	}
}

type responseEtl struct {
	Success bool `json:"success"`
}

type responseEtlError struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}
