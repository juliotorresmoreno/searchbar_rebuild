package db

import (
	"github.com/juliotorresmoreno/searchbar/config"
	mgo "gopkg.in/mgo.v2"
)

//MongoConnection Estructura con la conexion de la base de datos
type MongoConnection struct {
	*mgo.Session
	*mgo.Database
}

//GetMongoConnection Obtiene la conexion a la BD
func GetMongoConnection(source, database string) (MongoConnection, error) {
	var session, err = mgo.Dial(source)
	if err != nil {
		return MongoConnection{session, nil}, err
	}
	return MongoConnection{session, session.DB(database)}, nil
}

//GetMongoDb Obtiene la conexion a la BD
func GetMongoDb() (MongoConnection, error) {
	var session, err = mgo.Dial(config.MONGO_HOST + ":" + config.MONGO_PORT)
	if err != nil {
		return MongoConnection{session, nil}, err
	}
	return MongoConnection{session, session.DB(config.MONGO_DB)}, nil
}
