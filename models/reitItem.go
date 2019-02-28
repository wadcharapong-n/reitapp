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
	CeillingValue     string `bson:"ceillingValue"`
	ParNAV            string `bson:"parNAV"`
	Policy            string `bson:"policy"`
	PriceOfDay        string `bson:"priceOfDay"`
	MaxPriceOfDay     string `bson:"maxPriceOfDay"`
	MinPriceOfDay     string `bson:"minPriceOfDay"`
	NickName          string `bson:"nickName"`
}

type Favorite struct {
	// ID     string `bson:"_id"`
	Ticker string `bson:"ticker"`
	UserId string `bson:"userId"`
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
