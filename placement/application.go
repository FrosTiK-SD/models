package opportunity

import (
	"github.com/FrosTiK-SD/models-go/resume"
	"github.com/FrosTiK-SD/models-go/student"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Application struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Opportunity      primitive.ObjectID `json:"opportunity,omitempty"`
	Student          primitive.ObjectID `json:"student,omitempty"`
	Resume           primitive.ObjectID `json:"resume,omitempty"`
	DetailsRequested []DetailsRequested `json:"details_requested,omitempty"`
}

type ApplicationPopulated struct {
	ID               primitive.ObjectID   `json:"id,omitempty"`
	Opportunity      OpportunityPopulated `json:"opportunity,omitempty"`
	Student          student.Student      `json:"student,omitempty"`
	Resume           resume.Resume        `json:"resume,omitempty"`
	DetailsRequested []DetailsRequested   `json:"details_requested,omitempty"`
}

type DetailsRequested struct {
	Key   string `json:"key,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
