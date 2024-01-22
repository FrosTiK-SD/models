package company

import (
	constant "github.com/FrosTiK-SD/iitbhu-tpc-models-golang/constant"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JAF struct {
	ID               primitive.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty"`
	JAF_ID           primitive.ObjectID       `json:"jaf_id" bson:"jaf_id"`
	Recruiter        primitive.ObjectID       `json:"recruiter" bson:"recruiter"`
	Company          primitive.ObjectID       `json:"company" bson:"company"`
	Domain           string                   `json:"domain" bson:"domain"`
	AltHR            []AlternateHR            `json:"alternateHR" bson:"alternateHR"`
	JobDesc          JobDescription           `json:"jobDescription" bson:"jobDescription"`
	MedicalDetails   MedicalDetails           `json:"medical" bson:"medical"`
	SelectionProcess []string                 `json:"selectionProcess" bson:"selectionProcess"`
	ProcessMode      string                   `json:"processMode" bson:"processMode"`
	Session          constant.AcademicSession `json:"session" bson:"session"`
	IsActive         bool                     `json:"isActive" bson:"isActive"`
	RawKeyStore      map[string]interface{}   `json:"raw_key_store" bson:"raw_key_store"`
	CreatedAt        primitive.DateTime       `json:"createdAt" bson:"createdAt"`
	UpdatedAt        primitive.DateTime       `json:"updatedAt" bson:"updatedAt"`
}

type JAFPopulated struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	JAF_ID           primitive.ObjectID `json:"jaf_id" bson:"jaf_id"`
	Recruiter        Recruiter          `json:"recruiter" bson:"recruiter"`
	Company          Company            `json:"company" bson:"company"`
	AltHR            []AlternateHR      `json:"alternateHR" bson:"alternateHR"`
	JobDesc          JobDescription     `json:"jobDescription" bson:"jobDescription"`
	MedicalDetails   MedicalDetails     `json:"medical" bson:"medical"`
	SelectionProcess []string           `json:"selectionProcess" bson:"selectionProcess"`
	ProcessMode      string             `json:"processMode" bson:"processMode"`
	Session          string             `json:"session" bson:"session"`
	IsActive         bool               `json:"isActive" bson:"isActive"`
	CreatedAt        primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt        primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}

type JAFCSV struct {
	ID               string            `json:"_id" bson:"_id" csv:"ID"`
	JAF_ID           string            `json:"jaf_id" bson:"jaf_id" csv:"JAF ID"`
	Recruiter        RecruiterCSV      `json:"recruiter" bson:"recruiter" csv:"Recruiter ,inline"`
	Company          CompanyCSV        `json:"company" bson:"company" csv:"Company ,inline"`
	JobDesc          JobDescriptionCSV `json:"jobDescription" bson:"jobDescription" csv:",inline"`
	SelectionProcess string            `json:"selectionProcess" bson:"selectionProcess" csv:"Selection Process"`
	ProcessMode      string            `json:"processMode" bson:"processMode" csv:"Process Mode"`
	Session          string            `json:"session" bson:"session" csv:"Session"`
	IsActive         string            `json:"isActive" bson:"isActive" csv:"IsActive"`
	CreatedAt        string            `json:"createdAt" bson:"createdAt" csv:"Created At"`
	UpdatedAt        string            `json:"updatedAt" bson:"updatedAt" csv:"Updated At"`
}
