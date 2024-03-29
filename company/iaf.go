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

type IAFPopulated struct {
	ID               primitive.ObjectID    `json:"_id" bson:"_id"`
	IAF_ID           primitive.ObjectID    `json:"iaf_id" bson:"iaf_id"`
	Recruiter        Recruiter             `json:"recruiter" bson:"recruiter"`
	Company          Company               `json:"company" bson:"company"`
	AltHR            []AlternateHR         `json:"alternateHR" bson:"alternateHR"`
	InternshipDesc   InternshipDescription `json:"internshipDescription" bson:"internshipDescription"`
	MedicalDetails   MedicalDetails        `json:"medical" bson:"medical"`
	SelectionProcess []string              `json:"selectionProcess" bson:"selectionProcess"`
	ProcessMode      string                `json:"processMode" bson:"processMode"`
	Session          string                `json:"session" bson:"session"`
	IsActive         bool                  `json:"isActive" bson:"isActive"`
	CreatedAt        primitive.DateTime    `json:"createdAt" bson:"createdAt"`
	UpdatedAt        primitive.DateTime    `json:"updatedAt" bson:"updatedAt"`
}

type IAFCSV struct {
	ID               string                   `json:"_id" bson:"_id" csv:"ID"`
	IAF_ID           string                   `json:"iaf_id" bson:"iaf_id" csv:"IAF ID"`
	Recruiter        RecruiterCSV             `json:"recruiter" bson:"recruiter" csv:"Recruiter ,inline"`
	Company          CompanyCSV               `json:"company" bson:"company" csv:"Company ,inline"`
	InternshipDesc   InternshipDescriptionCSV `json:"internshipDescription" bson:"internshipDescription" csv:",inline"`
	SelectionProcess string                   `json:"selectionProcess" bson:"selectionProcess" csv:"Selection Process"`
	ProcessMode      string                   `json:"processMode" bson:"processMode" csv:"Process Mode"`
	Session          string                   `json:"session" bson:"session" csv:"Session"`
	IsActive         string                   `json:"isActive" bson:"isActive" csv:"IsActive"`
	CreatedAt        string                   `json:"createdAt" bson:"createdAt" csv:"Created At"`
	UpdatedAt        string                   `json:"updatedAt" bson:"updatedAt" csv:"Updated At"`
}
