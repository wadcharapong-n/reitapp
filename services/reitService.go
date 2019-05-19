package services

import (
	"encoding/json"
	"fmt"

	ai "github.com/night-codes/mgo-ai"
	"github.com/olivere/elastic"
	"github.com/spf13/viper"
	"github.com/wadcharapong/reitapp/app"
	"github.com/wadcharapong/reitapp/models"
	"github.com/wadcharapong/reitapp/util"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ReitServicer interface {
	GetReitBySymbol(symbol string) (models.ReitItem, error)
	GetReitAll() ([]models.ReitItem, error)
	SaveReitFavorite(userId string, symbol string) error
	DeleteReitFavorite(userId string, ticker string) error
	GetReitFavoriteByUserIDJoin(userId string) ([]*models.FavoriteInfo, error)
	GetUserProfileByCriteria(userId string, site string) models.UserProfile
	SaveUserProfile(profile *models.UserProfile) string
	CreateNewUserProfile(facebook models.Facebook, google models.Google) string
	SearchElastic(query string) []models.ReitItem
	SearchMap(lat float64, lon float64) models.PlaceInfo
	InsertReit(reit models.Reit) error
	UpdateUserProfile(profile *models.UserProfile) string
	GetUserAll() ([]models.UserProfile, error)
	DeleteUser(userId string) error
	GetPlaceAll() ([]models.Place, error)
	AddPlace(placeId string, name string, address string,
		symbol string, lat float64, long float64) error
	DeletePlace(placeId string) error
}

type Reit_Service struct {
	reitItems    []models.ReitItem
	reitItem     models.ReitItem
	reitFavorite []*models.FavoriteInfo
	userProfile  models.UserProfile
	userProfiles []models.UserProfile
	Place        models.Place
	Places       []models.Place
	locationInfo models.PlaceInfo
	err          error
	reit         []models.Reit
}

const reitCollection = "REIT"
const majorShareholdersCollection = "MajorShareholders"
const placeCollection = "Place"
const favoriteCollection = "Favorite"
const countersCollection = "counters"
const userProfileCollection = "UserProfile"

func (self Reit_Service) GetReitAll() ([]models.ReitItem, error) {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(reitCollection)
	query := []bson.M{{
		"$lookup": bson.M{ // lookup the documents table here
			"from":         majorShareholdersCollection,
			"localField":   "symbol",
			"foreignField": "symbol",
			"as":           "majorShareholders",
		}},
		{
			"$lookup": bson.M{ // lookup the documents table here
				"from":         placeCollection,
				"localField":   "symbol",
				"foreignField": "symbol",
				"as":           "place",
			}}}

	pipe := document.Pipe(query)
	self.err = pipe.All(&self.reitItems)
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.reitItems, self.err
}

func (self Reit_Service) GetReitBySymbol(symbol string) (models.ReitItem, error) {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(reitCollection)
	query := []bson.M{{
		"$lookup": bson.M{ // lookup the documents table here
			"from":         majorShareholdersCollection,
			"localField":   "symbol",
			"foreignField": "symbol",
			"as":           "majorShareholders",
		}},
		{
			"$lookup": bson.M{ // lookup the documents table here
				"from":         placeCollection,
				"localField":   "symbol",
				"foreignField": "symbol",
				"as":           "place",
			}},
		{"$match": bson.M{
			"symbol": symbol,
		}}}

	pipe := document.Pipe(query)
	self.err = pipe.One(&self.reitItem)
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.reitItem, self.err
}

func (self Reit_Service) SaveReitFavorite(userId string, symbol string) error {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	ai.Connect(session.DB(viper.GetString("mongodb.collection")).C(countersCollection))
	document := session.DB(viper.GetString("mongodb.collection")).C(favoriteCollection)
	favorite := models.Favorite{ID: ai.Next("Favorite"), UserId: userId, Symbol: symbol}
	self.err = document.Insert(&favorite)
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.err
}

func (self Reit_Service) DeleteReitFavorite(userId string, ticker string) error {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(favoriteCollection)
	self.err = document.Remove(bson.M{"symbol": ticker, "userId": userId})
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.err
}

func (self Reit_Service) GetReitFavoriteByUserIDJoin(userId string) ([]*models.FavoriteInfo, error) {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(favoriteCollection)

	query := []bson.M{{
		"$lookup": bson.M{ // lookup the documents table here
			"from":         reitCollection,
			"localField":   "symbol",
			"foreignField": "symbol",
			"as":           "Reit",
		}},
		{"$match": bson.M{
			"userId": userId,
		}}}

	pipe := document.Pipe(query)
	self.err = pipe.All(&self.reitFavorite)
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.reitFavorite, self.err
}

func (self Reit_Service) CreateNewUserProfile(facebook models.Facebook, google models.Google) string {
	var reitServicer ReitServicer
	reitServicer = Reit_Service{}
	var message string
	if (facebook != models.Facebook{}) {
		userProfile := reitServicer.GetUserProfileByCriteria(facebook.ID, "facebook")
		if (userProfile == models.UserProfile{}) {
			userProfile = models.UserProfile{
				UserID:   facebook.ID,
				UserName: facebook.Name,
				FullName: facebook.Name,
				Email:    facebook.Email,
				Image:    facebook.Picture.Data.URL,
				Site:     "facebook"}
			message = reitServicer.SaveUserProfile(&userProfile)
		} else {
			//case update
			userProfile.FullName = facebook.Name
			userProfile.UserName = facebook.Name
			userProfile.Email = facebook.Email
			userProfile.Image = facebook.Picture.Data.URL
			message = reitServicer.UpdateUserProfile(&userProfile)
		}
	} else if (google != models.Google{}) {
		userProfile := reitServicer.GetUserProfileByCriteria(google.ID, "google")
		if (userProfile == models.UserProfile{}) {
			userProfile = models.UserProfile{
				UserID:   google.ID,
				UserName: google.Name,
				FullName: google.Name,
				Image:    google.Picture,
				Email:    google.Email,
				Site:     "google"}
			message = reitServicer.SaveUserProfile(&userProfile)
		} else {
			//case update
			userProfile.UserName = facebook.Name
			userProfile.FullName = facebook.Name
			userProfile.Email = facebook.Email
			userProfile.Image = facebook.Picture.Data.URL
			message = reitServicer.UpdateUserProfile(&userProfile)
		}
	}
	return message
}

func (self Reit_Service) SaveUserProfile(profile *models.UserProfile) string {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	ai.Connect(session.DB(viper.GetString("mongodb.collection")).C(countersCollection))
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(userProfileCollection)
	profile.ID = ai.Next("userProfile")
	err := document.Insert(&profile)
	if err != nil {
		return util.FALI
	} else {
		return util.SUCCESS
	}
}

func (self Reit_Service) UpdateUserProfile(profile *models.UserProfile) string {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	ai.Connect(session.DB(viper.GetString("mongodb.collection")).C(countersCollection))
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(userProfileCollection)
	err := document.UpdateId(profile.ID, &profile)
	if err != nil {
		return util.FALI
	} else {
		return util.SUCCESS
	}
}

func (self Reit_Service) GetUserProfileByCriteria(userId string, site string) models.UserProfile {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(userProfileCollection)
	self.err = document.Find(bson.M{"userID": userId, "site": site}).One(&self.userProfile)
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.userProfile
}

func (self Reit_Service) SearchElastic(query string) []models.ReitItem {
	ctx := context.Background()
	client := app.GetElasticSearch()

	// Search with a term query
	termQuery := elastic.NewMultiMatchQuery(query, "nickName", "symbol", "reitManager").Type("phrase_prefix")
	//termQuery := elastic.NewTermQuery("nickName",query)
	searchResult, err := client.Search().
		Index(viper.GetString("elasticsearch.indexName")). // search in index "reitapp"
		Query(termQuery).                                  // specify the query
		//Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10). // take documents 0-9
		Pretty(true).     // pretty print request and response JSON
		Do(ctx)           // execute
	if err == nil {
		// TotalHits is another convenience function that works even when something goes wrong.
		fmt.Printf("Found a total of %d reits\n", searchResult.TotalHits())

		// Here's how you iterate through results with full control over each step.
		if searchResult.Hits.TotalHits > 0 {
			// Iterate through results
			for _, hit := range searchResult.Hits.Hits {
				// hit.Index contains the name of the index

				// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
				var t models.ReitItem
				err := json.Unmarshal(*hit.Source, &t)
				if err != nil {
					// Deserialization failed
				}
				self.reitItems = append(self.reitItems, t)
				// Work with tweet
				fmt.Printf("reit by %s: %s\n", t.Symbol, t.NickName)

			}
			return self.reitItems
		} else {
			// No hits
			fmt.Print("Found no reit\n")
		}
	} else {
		fmt.Printf("error : ", self.err)
	}
	return nil
}

func (self Reit_Service) SearchMap(lat float64, lon float64) models.PlaceInfo {
	fmt.Println("start : SearchMap")
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(placeCollection)
	scope := 100

	query := []bson.M{{
		"$geoNear": bson.M{
			"near": bson.M{
				"type":        "Point",
				"coordinates": []float64{lon, lat},
			},
			"distanceField": "dist.calculated",
			"maxDistance":   scope, // miles to meter
			"spherical":     "true",
		}},
		{"$lookup": bson.M{ // lookup the documents table here
			"from":         reitCollection,
			"localField":   "symbol",
			"foreignField": "symbol",
			"as":           "Reit",
		}}}

	pipe := document.Pipe(query)
	err := pipe.One(&self.locationInfo)
	if err != nil {
		fmt.Printf("error : ", err)
	}
	return self.locationInfo
}

func AddDataElastic(reit models.Reit) error {
	ctx := context.Background()
	client := app.GetElasticSearch()

	CheckIndex()
	//Search with a term query// Index a tweet (using JSON serialization)
	_, err := client.Index().
		Index(viper.GetString("elasticsearch.indexName")).
		Type("reit").
		//Id("1").
		BodyJson(&reit).
		Do(ctx)
	if err != nil {
		// Handle error
		fmt.Printf("error : ", err)
	}
	return err
}

func CheckIndex() {
	ctx := context.Background()
	client := app.GetElasticSearch()
	exists, err := client.IndexExists(viper.GetString("elasticsearch.indexName")).Do(ctx)
	if err == nil {
		if !exists {
			// Create a new index.
			createIndex, err := client.CreateIndex(viper.GetString("elasticsearch.indexName")).BodyString(models.Mapping).Do(ctx)
			if err != nil {
				// Handle error
				panic(err)
			}
			if !createIndex.Acknowledged {
				// Not acknowledged
			}
		}
	} else {
		fmt.Printf("error : ", err)
	}

}

func (self Reit_Service) GetReits() ([]models.Reit, error) {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(reitCollection)
	self.err = document.Find(bson.M{}).All(&self.reit)
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.reit, self.err
}

func (self Reit_Service) InsertReit(reit models.Reit) error {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	ai.Connect(session.DB(viper.GetString("mongodb.collection")).C(countersCollection))
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(reitCollection)
	reit.ID = ai.Next("reit")
	self.err = document.Insert(reit)
	if self.err == nil {
		self.err = AddDataElastic(reit)
	} else {
		fmt.Printf("error : ", self.err)
	}
	return self.err
}

func (self Reit_Service) GetUserAll() ([]models.UserProfile, error) {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(userProfileCollection)
	self.err = document.Find(nil).All(&self.userProfiles)
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.userProfiles, self.err

}

func (self Reit_Service) DeleteUser(userId string) error {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(userProfileCollection)
	self.err = document.Remove(bson.M{"userID": userId})
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.err

}

func (self Reit_Service) GetPlaceAll() ([]models.Place, error) {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(placeCollection)
	self.err = document.Find(nil).All(&self.Places)
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.Places, self.err

}

func (self Reit_Service) AddPlace(placeId string, name string, address string,
	symbol string, lat float64, long float64) error {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(placeCollection)
	self.err = document.Insert(bson.M{"placeId": placeId, "symbol": symbol, "name": name, "address": address, "location": bson.M{"type": "Point", "coordinates": []float64{long, lat}}})
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.err

}

func (self Reit_Service) DeletePlace(placeId string) error {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB(viper.GetString("mongodb.collection")).C(placeCollection)
	self.err = document.Remove(bson.M{"placeId": placeId})
	if self.err != nil {
		fmt.Printf("error : ", self.err)
	}
	return self.err

}
