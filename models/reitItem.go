package models

type ReitItem struct {
	ID                string `bson:"_id"`
	TrustNameTh       string `bson:"trustNameTh"`
	TrustNameEn       string `bson:"trustNameEn"`
	Ticker            string `bson:"ticker"`
	Trustee           string `bson:"trustee"`
	Address           string `bson:"address"`
	InvestmentAmount  string `bson:"investmentAmount"`
	EstablishmentDate string `bson:"establishmentDate"`
	RegistrationDate  string `bson:"registrationDate"`
	ReitManager       string `bson:"reitManager"`
}

type Favorite struct {
	// ID     string `bson:"_id"`
	Ticker string `bson:"ticker"`
	UserId string `bson:"userId"`
}
