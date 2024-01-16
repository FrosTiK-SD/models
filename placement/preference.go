package placement

import (
	"github.com/FrosTiK-SD/models-go/student"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Preference struct {
	Slot        primitive.ObjectID   `json:"slot,omitempty"`
	Student     primitive.ObjectID   `json:"student,omitempty"`
	Application []primitive.ObjectID `json:"application,omitempty"`
}

type PreferencePopulated struct {
	Slot        Slot            `json:"slot,omitempty"`
	Student     student.Student `json:"student,omitempty"`
	Application []Application   `json:"application,omitempty"`
}
