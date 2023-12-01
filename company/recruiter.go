package company

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recruiter struct {
	ID         primitive.ObjectID   `json:"_id" bson:"_id"`
	FirstName  string               `json:"firstName" bson:"firstName"`
	MiddleName string               `json:"middleName" bson:"middleName"`
	LastName   string               `json:"lastName" bson:"lastName"`
	Gender     string               `json:"gender" bson:"gender"`
	Contact    string               `json:"contact" bson:"contact"`
	AltContact string               `json:"alternateContact" bson:"alternateContact"`
	Email      string               `json:"email" bson:"email"`
	Company    primitive.ObjectID   `json:"company" bson:"company"`
	Groups     []primitive.ObjectID `json:"groups" bson:"groups"`
	IsActive   bool                 `json:"isActive" bson:"isActive"`
}

type RecruiterCSV struct {
	ID         string `json:"_id" bson:"_id" csv:"ID"`
	FullName   string `csv:"Full Name"`
	Gender     string `json:"gender" bson:"gender" csv:"Gender"`
	Contact    string `json:"contact" bson:"contact" csv:"Contact"`
	AltContact string `json:"alternateContact" bson:"alternateContact" csv:"Alt Contact"`
	Email      string `json:"email" bson:"email" csv:"Email"`
	Company    string `json:"company" bson:"company" csv:"Company"`
	IsActive   string `json:"is_active" bson:"is_active" csv:"IsActive"`
}
