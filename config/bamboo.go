package config

import (
	"os"

	"gopkg.in/mgo.v2"
)

//DB struct
type DB struct {
	Session *mgo.Session
}

//DoDial : create Database Connection
func (db *DB) DoDial() (s *mgo.Session, err error) {
	return mgo.Dial(DBUrl())
}

//Name : Database Name
func (db *DB) Name() string {
	return "bambooshade"
}

//DBUrl ... setting Database URI
func DBUrl() string {
	dburl := os.Getenv("MONGODB_URL")

	if dburl == "" {
		dburl = "localhost"
	}

	return dburl
}
