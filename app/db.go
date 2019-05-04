package app

import (
	"github.com/olivere/elastic"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

func GetDocumentMongo() *mgo.Session {
	session, err := mgo.Dial(viper.GetString("mongodb.connection"))
	if err != nil {
		panic(err)
	}
	return session
}

func GetElasticSearch() *elastic.Client {

	elasticClient, err := elastic.NewClient(
		elastic.SetURL(viper.GetString("elasticsearch.connection")),
		elastic.SetSniff(false),
	)
	if err != nil {
		// Handle error
		panic(err)
	}

	return elasticClient
}
