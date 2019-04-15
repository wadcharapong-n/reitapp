package models

import (
	"github.com/dgrijalva/jwt-go"
)

type ReitItem struct {
	ID                string `json:"-" bson:"_id"`
	TrustNameTh       string `json:"trustNameTh" bson:"trustNameTh"`
	TrustNameEn       string `json:"trustNameEn" bson:"trustNameEn"`
	Symbol            string `json:"symbol" bson:"symbol"`
	Trustee           string `json:"trustee" bson:"trustee"`
	Address           string `json:"address" bson:"address"`
	InvestmentAmount  string `json:"investmentAmount" bson:"investmentAmount"`
	EstablishmentDate string `json:"establishmentDate" bson:"establishmentDate"`
	RegistrationDate  string `json:"registrationDate" bson:"registrationDate"`
	ReitManager       string `json:"reitManager" bson:"reitManager"`
	ParValue          string `json:"parValue" bson:"parValue"`
	CeilingValue      string `json:"ceilingValue" bson:"ceilingValue"`
	FloorValue		  string `json:"floorValue" bson:"floorValue"`
	PeValue           string `json:"parNAV" bson:"parNAV"`
	ParNAV            string `json:"peValue" bson:"peValue"`
	Policy            string `json:"policy" bson:"policy"`
	PriceOfDay        string `json:"priceOfDay" bson:"priceOfDay"`
	MaxPriceOfDay     string `json:"maxPriceOfDay" bson:"maxPriceOfDay"`
	MinPriceOfDay     string `json:"minPriceOfDay" bson:"minPriceOfDay"`
	NickName          string `json:"nickName" bson:"nickName"`
	Location 		  GeoJson `json:"location" bson:"location"`
	MajorShareholders []MajorShareholders `json:"majorShareholders" bson:"majorShareholders"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates [][]float64 `json:"coordinates"`
}

type Favorite struct {
	ID     string `json:"-" bson:"_id"`
	Symbol string `json:"symbol" bson:"symbol"`
	UserId string `json:"userId" bson:"userId"`
}

type FavoriteInfo struct {
	// ID     string `bson:"_id"`
	Symbol string `json:"symbol" bson:"symbol"`
	UserId string `json:"userId" bson:"userId"`
	ReitItem []ReitItem `json:"Reit" bson:"Reit"`
}

type MajorShareholders struct {
	ID     string `json:"-" bson:"_id"`
	Symbol 		string `json:"symbol" bson:"symbol"`
	NameTh      string `json:"nameTh" bson:"nameTh"`
	NameEn      string `json:"nameEn" bson:"nameEn"`
	Shares		string `json:"shares" bson:"shares"`
	SharesPercent		string `json:"sharesPercent" bson:"sharesPercent"`
}

type JWTCustomClaims struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
	Site string `bson:"site"`
	jwt.StandardClaims
}

type UserProfile struct {
	//ID       string `bson:"_id"`
	UserID   string `json:"userID" bson:"userID"`
	UserName string `json:"userName" bson:"userName"`
	FullName string `json:"fullName" bson:"fullName"`
	Email    string `json:"email" bson:"email"`
	Image    string `json:"image" bson:"image"`
	Site     string `json:"site" bson:"site"`
}

type Facebook struct {
	ID     	string `bson:"id"`
	Name 	string `bson:"name"`
	Email	string `bson:"email"`
}

type Google struct {
	ID     				string `bson:"id"`
	Name 				string `bson:"name"`
	Email 				string `bson:"email"`
	Verified_Email 		bool 	`bson:"verified_email"`
	Given_Name 			string `bson:"given_name"`
	Family_Name 		string `bson:"family_name"`
	Link 				string `bson:"link"`
	Picture 			string `bson:"picture"`
	Gender 				string `bson:"gender"`
	Locale 				string `bson:"locale"`
}

const Mapping  =`
    {
	"mappings" : {
      "reit" : {
        "properties" : {
          "Address" : {
            "type" : "text",
			"analyzer": "standard"
          },
          "ID" : {
            "type" : "text",
            "fields" : {
              "keyword" : {
                "type" : "keyword",
                "ignore_above" : 256
              }
            }
          },
          "InvestmentAmount" : {
            "type" : "text",
			"analyzer": "standard",
            "fields" : {
              "keyword" : {
                "type" : "keyword",
                "ignore_above" : 256
              }
            }
          },
          "NickName" : {
            "type" : "text",
			"analyzer": "standard",
            "fields" : {
              "keyword" : {
                "type" : "keyword",
                "ignore_above" : 256
              }
            }
          },
          "ReitManager" : {
            "type" : "text",
            "fields" : {
              "keyword" : {
                "type" : "keyword",
                "ignore_above" : 256
              }
            }
          },
          "Symbol" : {
            "type" : "text",
            "analyzer": "standard"
          },
          "TrustNameEn" : {
            "type" : "text",
			"analyzer": "standard"
          },
          "TrustNameTh" : {
            "type" : "text",
            "analyzer": "standard"
          },
          "Trustee" : {
            "type" : "text",
            "analyzer": "standard"
          }
        }
      }
    }
	}`
