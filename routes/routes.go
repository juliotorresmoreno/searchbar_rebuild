package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/searchbar/controllers"
	"github.com/nytimes/gziphandler"
)

//GetRoutes Obtiene el enrutador del sistema
func GetRoutes() *mux.Router {
	var mux = mux.NewRouter().StrictSlash(false)
	mux.HandleFunc("/import", controllers.GetEtl).Methods("GET")
	mux.HandleFunc("/query", controllers.GetQuery).Methods("GET")
	mux.PathPrefix("/").HandlerFunc(controllers.GetIndex).Methods("GET")
	withoutGz := http.Handler(mux)
	withGz := gziphandler.GzipHandler(withoutGz)
	mux.PathPrefix("/").Handler(withGz).Methods("GET")
	return mux
}
