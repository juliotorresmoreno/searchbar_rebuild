package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

//PORT puerto de escucha del servidor http
var PORT string

//REDIS_HOST servidor de base de datos redis
var REDIS_HOST string

//REDIS_PORT Puerto de la base de datos redis
var REDIS_PORT string

//SOURCES Listado de origenes de datos
var SOURCES []source

//READ_TIMEOUT Tiempo de espera que tardara el servidor en recibir los datos
var READ_TIMEOUT time.Duration

func init() {
	file := "./config/config.json"
	text, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var data = &configuration{}
	err = json.Unmarshal(text, data)
	if err != nil {
		panic(err)
	}
	PORT = data.Port
	READ_TIMEOUT = data.ReadTimeout
	REDIS_HOST = data.RedisHost
	REDIS_PORT = data.RedisPort
	SOURCES = data.Sources
}
