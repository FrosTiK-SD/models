package opportunity

import (
	"github.com/FrosTiK-SD/models/misc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileApplication struct {
	Resume                 string                  `json:"resume" bson:"resume"`
	Profile                primitive.ObjectID      `json:"profile" bson:"profile"`
	DetailsRequestedSchema *map[string]interface{} `bson:"detailsRequestedSchema" json:"detailsRequestedSchema"`
	Activities             []misc.Activity         `json:"activities" bson:"activities"`
}

type Application struct {
	ID                     primitive.ObjectID      `bson:"_id" json:"_id,omitempty"`
	Opportunity            primitive.ObjectID      `bson:"opportunity" json:"opportunity,omitempty"`
	Student                primitive.ObjectID      `bson:"student" json:"student,omitempty"`
	DetailsRequestedSchema *map[string]interface{} `bson:"detailsRequestedSchema" json:"detailsRequestedSchema"`
	Profiles               []ProfileApplication    `json:"profiles" bson:"profiles"`
}
