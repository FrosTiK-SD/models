package opportunity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileApplication struct {
	Resume     string               `json:"resume" bson:"resume"` // Link to the resume
	Profile    primitive.ObjectID   `json:"profile" bson:"profile"`
	Activities []primitive.ObjectID `json:"activities" bson:"activities"`

	// This is the submitted detais of the user on an CompanyProfile basis
	DetailsRequested *map[string]interface{} `json:"detailsRequested" bson:"detailsRequested"`
}

type Application struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Opportunity primitive.ObjectID `bson:"opportunity" json:"opportunity,omitempty"`
	Student     primitive.ObjectID `bson:"student" json:"student,omitempty"`

	// This is the submitted detais of the user on an Opportunity basis
	DetailsRequested *map[string]interface{} `json:"detailsRequested" bson:"detailsRequested"`
	Profiles         []ProfileApplication    `json:"profiles" bson:"profiles"`
}
