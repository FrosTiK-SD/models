package opportunity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Preference struct {
	Slot        primitive.ObjectID   `json:"slot,omitempty"`
	Student     primitive.ObjectID   `json:"student,omitempty"`
	Application []primitive.ObjectID `json:"application,omitempty"`
}
