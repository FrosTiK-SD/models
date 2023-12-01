package company

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OID struct {
	ID primitive.ObjectID `json:"id"`
}

type AlternateHR struct {
	Name       string `json:"name" bson:"name"`
	Contact    string `json:"contact" bson:"contact"`
	Email      string `json:"email" bson:"email"`
	AltContact string `json:"alternativeContact" bson:"alternativeContact"`
}

type EligibilityCriteria struct {
	CGPA     float32  `json:"cgpa" bson:"cgpa"`
	Branches []string `json:"branches" bson:"branches"`
	AgeLimit string   `json:"ageLimit" bson:"ageLimit"`
}

type EligibilityCriteriaCSV struct {
	CGPA     string `json:"cgpa" bson:"cgpa" csv:"CGPA"`
	Branches string `json:"branches" bson:"branches" csv:"Branches"`
	AgeLimit string `json:"ageLimit" bson:"ageLimit" csv:"Age Limit"`
}

type MedicalDetails struct {
	ColourBlindness string `json:"colorBlindness" bson:"colorBlindness"`
	Visibility      string `json:"visibility" bson:"visibility"`
	Height          string `json:"height" bson:"height"`
	Weight          string `json:"weight" bson:"weight"`
	Others          string `json:"others" bson:"others"`
}

type DBRef struct {
	Ref string             `bson:"$ref" json:"$ref"`
	ID  primitive.ObjectID `bson:"$id" json:"$id"`
	DB  string             `bson:"$db" json:"$db"`
}

type StudentApplication struct {
	StudentID  DBRef
	ResumeID   DBRef
	AppliedFor DBRef
}

type Group struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Roles []string           `json:"roles" bson:"roles"`
}

type AuthDetails struct {
	Email            string   `json:"email" bson:"email"`
	CompaniesAlloted []string `json:"companiesAlloted" bson:"companiesAlloted"`
	Groups           []Group  `json:"groups" bson:"groups"`
}
