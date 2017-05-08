package controllers

import (
	"fmt"
	"net/http"
)

//GetIndex Obtiene la pagina de inicio de la aplicacion
func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola mundo")
}
