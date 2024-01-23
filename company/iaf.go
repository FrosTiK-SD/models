package company

import (
	constant "github.com/FrosTiK-SD/models/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAF struct {
	ID               primitive.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty"`
	IAF_ID           primitive.ObjectID       `json:"iaf_id" bson:"iaf_id"`
	Recruiter        primitive.ObjectID       `json:"recruiter" bson:"recruiter"`
	Company          primitive.ObjectID       `json:"company" bson:"company"`
	Domain           string                   `json:"domain" bson:"domain"`
	AltHR            []AlternateHR            `json:"alternateHR" bson:"alternateHR"`
	InternshipDesc   InternshipDescription    `json:"internshipDescription" bson:"internshipDescription"`
	MedicalDetails   MedicalDetails           `json:"medical" bson:"medical"`
	SelectionProcess []string                 `json:"selectionProcess" bson:"selectionProcess"`
	ProcessMode      string                   `json:"processMode" bson:"processMode"`
	Session          constant.AcademicSession `json:"session" bson:"session"`
	IsActive         bool                     `json:"isActive" bson:"isActive"`
	RawKeyStore      map[string]interface{}   `json:"raw_key_store" bson:"raw_key_store"`
	CreatedAt        primitive.DateTime       `json:"createdAt" bson:"createdAt"`
	UpdatedAt        primitive.DateTime       `json:"updatedAt" bson:"updatedAt"`
}
