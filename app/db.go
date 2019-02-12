package app

import (
	"gopkg.in/mgo.v2"
)

func GetDocumentMongo() *mgo.Session {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	return session
}
