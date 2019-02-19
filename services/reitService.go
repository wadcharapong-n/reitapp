package services

import (
	"../app"
	"../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetReitAll() ([]*models.ReitItem, error) {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("REIT")
	results := []*models.ReitItem{}
	err := document.Find(nil).All(&results)
	return results, err
}

func GetReitBySymbol(symbol string) (models.ReitItem, error) {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("REIT")
	result := models.ReitItem{}
	err := document.Find(bson.M{"symbol": symbol}).One(&result)
	return result, err
}
