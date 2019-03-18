package services

import (
	"../app"
	"../models"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ReitServicer interface {
	GetReitBySymbol(symbol string) (models.ReitItem, error)
	GetReitAll() ([]*models.ReitItem, error)
	SaveReitFavorite(userId string, symbol string) error
	DeleteReitFavorite(userId string, ticker string) error
	GetReitFavoriteByUserIDJoin(userId string) []*models.FavoriteInfo
	GetUserProfileByCriteria(userId string, site string ) models.UserProfile
	SaveUserProfile(profile *models.UserProfile) string
}

type Reit_Service struct {
	reitItems []*models.ReitItem
	reitItem models.ReitItem
	reitFavorite []*models.FavoriteInfo
	userProfile models.UserProfile
	err error
}

func GetReitAllProcess(reitService ReitServicer) ([]*models.ReitItem, error) {
	return reitService.GetReitAll()
}

func GetReitBySymbolProcess(reitService ReitServicer,symbol string) (models.ReitItem, error) {
	return reitService.GetReitBySymbol(symbol)
}

func GetReitFavoriteByUserIDJoinProcess(reitService ReitServicer,userId string) []*models.FavoriteInfo {
	return reitService.GetReitFavoriteByUserIDJoin(userId)
}

func SaveReitFavoriteProcess(reitService ReitServicer,userId string, symbol string) error {
	return reitService.SaveReitFavorite(userId,symbol)
}

func GetUserProfileByCriteriaProcess(reitService ReitServicer,userId string, site string) models.UserProfile {
	return reitService.GetUserProfileByCriteria(userId,site)
}

func DeleteReitFavoriteProcess(reitService ReitServicer,userId string, symbol string) error{
	return reitService.DeleteReitFavorite(userId,symbol)
}

func SaveUserProfileProcess(reitService ReitServicer,profile *models.UserProfile) string {
	return reitService.SaveUserProfile(profile);
}

func (self Reit_Service) GetReitAll() ([]*models.ReitItem, error) {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("REIT")
	self.err = document.Find(nil).All(&self.reitItems)
	return self.reitItems, self.err
}


func (self Reit_Service) GetReitBySymbol(symbol string) (models.ReitItem, error) {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("REIT")
	self.err = document.Find(bson.M{"symbol": symbol}).One(&self.reitItem)
	return self.reitItem, self.err
}



func (self Reit_Service) SaveReitFavorite(userId string, symbol string) error {
	fmt.Println("start : GetReitAll")
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("Favorite")
	favorite := models.Favorite{UserId: userId, Symbol: symbol}
	self.err = document.Insert(&favorite)
	if self.err != nil {
		// TODO: Do something about the error
		fmt.Printf("error : ", self.err)
	} else {

	}
	return self.err
}

func (self Reit_Service) DeleteReitFavorite(userId string, ticker string) error{
	fmt.Println("start : GetReitAll")
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("Favorite")
	favorite := models.Favorite{UserId: userId, Symbol: ticker}
	self.err = document.Remove(&favorite)
	if self.err != nil {
		// TODO: Do something about the error
		fmt.Printf("error : ", self.err)
	} else {

	}
	return self.err
}

//func GetReitFavoriteByUserID(userId string) []*models.Favorite {
//	fmt.Println("start : GetReitAll")
//	session := *app.GetDocumentMongo()
//	defer session.Close()
//	// Optional. Switch the session to a monotonic behavior.
//	session.SetMode(mgo.Monotonic, true)
//	document := session.DB("REIT_DEV").C("Favorite")
//	results := []*models.Favorite{}
//
//	err := document.Find(bson.M{"userId": userId }).All(&results)
//	if err != nil {
//		// TODO: Do something about the error
//		fmt.Printf("error : ", err)
//	} else {
//
//	}
//	return results
//}

func (self Reit_Service) GetReitFavoriteByUserIDJoin(userId string) []*models.FavoriteInfo {
	fmt.Println("start : GetReitAll")
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("Favorite")

	query := []bson.M{{
		"$lookup": bson.M{ // lookup the documents table here
			"from":         "REIT",
			"localField":   "symbol",
			"foreignField": "symbol",
			"as":           "Reit",
		}},
		{"$match": bson.M{
			"userId": userId,
		}}}

	pipe := document.Pipe(query)
	err := pipe.All(&self.reitFavorite)
	if err != nil {
		// TODO: Do something about the error
		fmt.Printf("error : ", err)
	} else {

	}
	return self.reitFavorite
}

func CreateNewUserProfile(facebook models.Facebook,google models.Google ) string {
	var reitServicer ReitServicer
	reitServicer = Reit_Service{}
	var message string
	if (facebook != models.Facebook{}){
		userProfile := reitServicer.GetUserProfileByCriteria(facebook.ID, "facebook")
		if(userProfile == models.UserProfile{}){
			userProfile =  models.UserProfile{
				UserID: facebook.ID,
				UserName: facebook.Name,
				FullName:facebook.Name,
				Email:facebook.Email,
				Site:"facebook"}
			message = reitServicer.SaveUserProfile(&userProfile)
		}
	} else if (google != models.Google{}){
		userProfile := reitServicer.GetUserProfileByCriteria(google.ID, "google")
		if(userProfile == models.UserProfile{}){
			userProfile =  models.UserProfile{
				UserID: google.ID,
				UserName: google.Name,
				FullName:google.Name,
				Image:google.Picture,
				Email:google.Email,
				Site:"google"}
			message = reitServicer.SaveUserProfile(&userProfile)
		}
	}
	return message
}

func (self Reit_Service) SaveUserProfile(profile *models.UserProfile) string {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("UserProfile")
	err := document.Insert(&profile)
	if err != nil {
		return "fail"
	} else {
		return "success"
	}
}

func (self Reit_Service) GetUserProfileByCriteria(userId string, site string ) models.UserProfile {
	session := *app.GetDocumentMongo()
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	document := session.DB("REIT_DEV").C("UserProfile")
	err := document.Find(bson.M{"userID": userId,"site": site}).One(&self.userProfile)
	if err != nil {
		// TODO: Do something about the error
		fmt.Printf("error : ", err)
	} else {

	}
	return self.userProfile
}
