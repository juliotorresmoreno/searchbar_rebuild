package bootstrap

import (
	"net/http"
	"time"
	"github.com/juliotorresmoreno/searchbar/config"
	"github.com/juliotorresmoreno/searchbar/routes"
)

func StartHTTP() {
	var mux = routes.GetRoutes()
	var addr = ":" + config.PORT
	var server = &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    config.READ_TIMEOUT * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	println("Listening on " + addr)
	println(server.ListenAndServe())
}