package opportunity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Application struct {
	ID                     primitive.ObjectID     `bson:"_id" json:"_id,omitempty"`
	Opportunity            primitive.ObjectID     `bson:"opportunity" json:"opportunity,omitempty"`
	Student                primitive.ObjectID     `bson:"student" json:"student,omitempty"`
	Resume                 primitive.ObjectID     `bson:"resume" json:"resume,omitempty"`
	DetailsRequestedSchema map[string]interface{} `bson:"detailsRequestedSchema" json:"detailsRequestedSchema"`
}
