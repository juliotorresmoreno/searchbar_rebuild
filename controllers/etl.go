package controllers

import (
	"fmt"
	"net/http"

	"github.com/juliotorresmoreno/searchbar/config"
	"github.com/juliotorresmoreno/searchbar/etl"
)

//GetEtl Metodo Get encargado de importar los datos
func GetEtl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Importando data\n")
	tool := etl.NewEtl()
	tool.Open()
	defer tool.Close()
	for _, v := range config.SOURCES {
		err := tool.ImportData(v.Host, v.Database, v.Collection)
		if err != nil {
			fmt.Fprint(w, err.Error()+"\n")
		} else {
			fmt.Fprint(w, "Exito\n")
		}
	}
	fmt.Fprint(w, "Importacion completada")
}
