package db

import mgo "gopkg.in/mgo.v2"

// GetConnection
func GetMongoConnection(source, database string) (*mgo.Session, *mgo.Database, error) {
	var session, err = mgo.Dial(source)
	if err != nil {
		return session, nil, err
	}
	return session, session.DB(database), nil
}