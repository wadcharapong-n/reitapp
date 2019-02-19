package services

import (
	"../app"
	"../models"
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
)

func GetReitAll(c echo.Context) []*models.ReitItem {
	fmt.Println("start : GetReitAll")
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

func SearchElastic() string {
	client := app.GetElasticSearch()
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
	return esversion
}
