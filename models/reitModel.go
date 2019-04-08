package models

import "github.com/dgrijalva/jwt-go"

type ReitItem struct {
	ID                string `bson:"_id"`
	TrustNameTh       string `bson:"trustNameTh"`
	TrustNameEn       string `bson:"trustNameEn"`
	Symbol            string `bson:"symbol"`
	Trustee           string `bson:"trustee"`
	Address           string `bson:"address"`
	InvestmentAmount  string `bson:"investmentAmount"`
	EstablishmentDate string `bson:"establishmentDate"`
	RegistrationDate  string `bson:"registrationDate"`
	ReitManager       string `bson:"reitManager"`
	ParValue          string `bson:"parValue"`
	CeilingValue      string `bson:"ceilingValue"`
	FloorValue		  string `bson:"floorValue"`
	PeValue           string `bson:"parNAV"`
	ParNAV            string `bson:"peValue"`
	Policy            string `bson:"policy"`
	PriceOfDay        string `bson:"priceOfDay"`
	MaxPriceOfDay     string `bson:"maxPriceOfDay"`
	MinPriceOfDay     string `bson:"minPriceOfDay"`
	NickName          string `bson:"nickName"`
}

type Favorite struct {
	// ID     string `bson:"_id"`
	Symbol string `bson:"symbol"`
	UserId string `bson:"userId"`
}

type FavoriteInfo struct {
	// ID     string `bson:"_id"`
	Symbol string `bson:"symbol"`
	UserId string `bson:"userId"`
	ReitItem []ReitItem `bson:"Reit"`
}

type JWTCustomClaims struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
	Site string `bson:"site"`
	jwt.StandardClaims
}

type UserProfile struct {
	//ID       string `bson:"_id"`
	UserID   string `bson:"userID"`
	UserName string `bson:"userName"`
	FullName string `bson:"fullName"`
	Email    string `bson:"email"`
	Image    string `bson:"image"`
	Site     string `bson:"site"`
}

type Facebook struct {
	ID     string `bson:"id"`
	Name string `bson:"name"`
	Email string `bson:"email"`
}

type Google struct {
	ID     string `bson:"id"`
	Name string `bson:"name"`
	Email string `bson:"email"`
	Verified_Email bool `bson:"verified_email"`
	Given_Name string `bson:"given_name"`
	Family_Name string `bson:"family_name"`
	Link string `bson:"link"`
	Picture string `bson:"picture"`
	Gender string `bson:"gender"`
	Locale string `bson:"locale"`
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
