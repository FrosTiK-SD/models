package company

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	PRIVATE = "private"
	PSU     = "Public Sector Unit"
)

type Company struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	LogoURLs        []string           `json:"logo" bson:"logo"`
	Website         string             `json:"website" bson:"website"`
	Address         string             `json:"address" bson:"address"`
	Category        string             `json:"category" bson:"category"`
	Sector          string             `json:"sector" bson:"sector"`
	CompanyTurnover string             `json:"companyTurnover" bson:"companyTurnover"`
}

type CompanyCSV struct {
	ID              string `json:"_id" bson:"_id" csv:"ID"`
	Name            string `json:"name" bson:"name" csv:"Name"`
	Website         string `json:"website" bson:"website" csv:"Website"`
	Address         string `json:"address" bson:"address" csv:"Address"`
	Category        string `json:"category" bson:"category" csv:"Category"`
	Sector          string `json:"sector" bson:"sector" csv:"Sector"`
	CompanyTurnover string `json:"companyTurnover" bson:"companyTurnover" csv:"Turnover"`
}
