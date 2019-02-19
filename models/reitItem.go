package models

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
