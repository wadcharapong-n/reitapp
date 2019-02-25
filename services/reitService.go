package services

import (
	"../app"
	"../config"
	"../models"
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/olivere/elastic.v6"
	"log"
	"net/http"
	"strconv"
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

func SearchElastic(e echo.Context) error {
	client := app.GetElasticSearch()
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Parse request
	query := e.QueryParam("query")
	if query == "" {
		e.JSON(http.StatusBadRequest, "Query not specified")
		return nil
	}
	skip := 0
	take := 10
	if i, err := strconv.Atoi( e.QueryParam("skip")); err == nil {
		skip = i
	}
	if i, err := strconv.Atoi( e.QueryParam("take")); err == nil {
		take = i
	}
	// Perform search
	esQuery := elastic.NewMultiMatchQuery(query, "title", "content").
		Fuzziness("2").
		MinimumShouldMatch("2")
	result, err := client.Search().
		Index(config.ElasticIndexName).
		Query(esQuery).
		From(skip).Size(take).
		Do(e.Request().Context())
	if err != nil {
		log.Println(err)
		e.JSON(http.StatusInternalServerError, "Something went wrong")
		return nil
	}
	//prepare response from elasticsearch
	//res := SearchResponse{
	//	Time: fmt.Sprintf("%d", result.TookInMillis),
	//	Hits: fmt.Sprintf("%d", result.Hits.TotalHits),
	//}
	//// Transform search results before returning them
	//docs := make([]DocumentResponse, 0)
	//for _, hit := range result.Hits.Hits {
	//	var doc DocumentResponse
	//	json.Unmarshal(*hit.Source, &doc)
	//	docs = append(docs, doc)
	//}
	//res.Documents = docs
	e.JSON(http.StatusOK, result)

	return nil
}
