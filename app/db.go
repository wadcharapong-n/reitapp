package app

import (
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	"github.com/wadcharapong/reitapp/config"
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
