package app

import (
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
)

func GetDocumentMongo() *mgo.Session {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	return session
}

func GetElasticSearch() *elastic.Client {
	elasticClient, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		// Handle error
		panic(err)
	}

	return elasticClient
}
