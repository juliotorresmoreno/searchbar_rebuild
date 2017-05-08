package db

import mgo "gopkg.in/mgo.v2"

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
