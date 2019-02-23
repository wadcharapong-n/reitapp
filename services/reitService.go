package services

import (
	"../app"
	"../models"
	"fmt"
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

func SaveReitFavorite(userId string, ticker string) {
	fmt.Println("start : GetReitAll")
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("Favorite")
	favorite := models.Favorite{UserId: userId, Ticker: ticker}
	err := document.Insert(&favorite)
	if err != nil {
		// TODO: Do something about the error
		fmt.Printf("error : ", err)
	} else {

	}
}

func GetReitFavoriteByUserID(userId string) []*models.Favorite {
	fmt.Println("start : GetReitAll")
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("Favorite")
	results := []*models.Favorite{}
	err := document.Find(bson.M{"userId": userId}).All(&results)
	if err != nil {
		// TODO: Do something about the error
		fmt.Printf("error : ", err)
	} else {

	}
	return results
}
