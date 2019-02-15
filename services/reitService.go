package services

import (
	"../app"
	"../models"
	"fmt"
	"gopkg.in/mgo.v2"
)

func GetReitAll() []*models.ReitItem {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("REIT")
	results := []*models.ReitItem{}
	err := document.Find(nil).All(&results)
	if err != nil {
		// TODO: Do something about the error
		fmt.Printf("error : ", err)
	} else {

	}
	return results
}
