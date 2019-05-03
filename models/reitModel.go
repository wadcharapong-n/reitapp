package models

import (
	"github.com/dgrijalva/jwt-go"
)

type ReitItem struct {
	ID     			  uint64 `json:"-" bson:"_id"`
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
	URL				  string `json:"url" bson:"url"`
	PropertyManager string `json:"propertyManager" bson:"propertyManager"`
	InvestmentPolicy string `json:"investmentPolicy" bson:"investmentPolicy"`
	MajorShareholders []MajorShareholders `json:"majorShareholders" bson:"majorShareholders"`
	Place []Place `json:"place" bson:"place"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

type Place struct{
	ID     uint64 `json:"-" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	Symbol string `json:"symbol" bson:"symbol"`
	Location GeoJson `json:"location" bson:"location"`
}

type PlaceInfo struct{
	ID     uint64 `json:"-" bson:"_id"`
	Symbol string `json:"symbol" bson:"symbol"`
	Name string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	Location GeoJson `json:"location" bson:"location"`
	ReitItem []ReitItem `json:"Reit" bson:"Reit"`
}

type Favorite struct {
	ID     uint64 `json:"-" bson:"_id"`
	Symbol string `json:"symbol" bson:"symbol"`
	UserId string `json:"userId" bson:"userId"`
}

type FavoriteInfo struct {
	ID     uint64 `json:"-" bson:"_id"`
	Symbol string `json:"symbol" bson:"symbol"`
	UserId string `json:"userId" bson:"userId"`
	ReitItem []ReitItem `json:"Reit" bson:"Reit"`
}

type MajorShareholders struct {
	ID     uint64 `json:"-" bson:"_id"`
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
	ID     uint64 `json:"-" bson:"_id"`
	UserID   string `json:"userID" bson:"userID"`
	UserName string `json:"userName" bson:"userName"`
	FullName string `json:"fullName" bson:"fullName"`
	Email    string `json:"email" bson:"email"`
	Image    string `json:"image" bson:"image"`
	Site     string `json:"site" bson:"site"`
}

type Facebook struct {
	ID     	string `json:"id"`
	Name 	string `json:"name"`
	Email	string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Picture Picture `json:"picture"`
}

type Picture struct {
	Data Data `json:"data"`
}

type Data struct {
	Height     		string `json:"height"`
	IsSilhouette    string `json:"is_silhouette"`
	URL     		string `json:"url"`
	Width     		string `json:"width"`
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
