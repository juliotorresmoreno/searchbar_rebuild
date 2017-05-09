package controllers

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/juliotorresmoreno/searchbar/lib"
)

//GetGraphQL Metodo Get encargado de describir y actualizar los datos
func GetGraphQL(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["query"]
	if len(query) == 0 {
		fmt.Fprint(w, "{\"success\": false, \"error\": \"Mal formado\"}")
		return
	}
	result := lib.ExecuteQuery(query[0])
	json.NewEncoder(w).Encode(result)
}
