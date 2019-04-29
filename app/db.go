package app

import (
	"github.com/olivere/elastic"
	"github.com/wadcharapong/reitapp/config"
	"gopkg.in/mgo.v2"
)

func GetDocumentMongo() *mgo.Session {
	session, err := mgo.Dial(config.Mongo_URL)
	if err != nil {
		panic(err)
	}
	return session
}

func GetElasticSearch() *elastic.Client {

	elasticClient, err := elastic.NewClient(
		elastic.SetURL(config.Elastic_URL),
		elastic.SetSniff(false),
	)
	if err != nil {
		// Handle error
		panic(err)
	}

	return elasticClient
}
