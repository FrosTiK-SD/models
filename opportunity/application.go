package opportunity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Application struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Opportunity      primitive.ObjectID `json:"opportunity,omitempty"`
	Student          primitive.ObjectID `json:"student,omitempty"`
	Resume           primitive.ObjectID `json:"resume,omitempty"`
	DetailsRequested []DetailsRequested `json:"details_requested,omitempty"`
}

type DetailsRequested struct {
	Key   string `json:"key,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
