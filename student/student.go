package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	group "github.com/FrosTiK-SD/models/group"
)

type Student struct {
	ID               primitive.ObjectID   `json:"_id" bson:"_id"`
	Groups           []primitive.ObjectID `json:"groups" bson:"groups"`
	CompaniesAlloted []string             `json:"companiesAlloted" bson:"companiesAlloted"`
	Batch            int                  `json:"batch" bson:"batch"`
	RollNo           int                  `json:"rollNo" bson:"rollNo"`
	FirstName        string               `json:"firstName" bson:"firstName"`
	LastName         string               `json:"lastName" bson:"lastName"`
	Department       string               `json:"department" bson:"department"`
	Course           string               `json:"course" bson:"course"`
	Email            string               `json:"email" bson:"email"`
	PersonalEmail    string               `json:"personalEmail" bson:"personalEmail"`
	LinkedIn         string               `json:"linkedIn" bson:"linkedIn"`
	Github           string               `json:"github" bson:"github"`
	MicrosoftTeams   string               `json:"microsoftTeams" bson:"microsoftTeams"`
	Mobile           int64                `json:"mobile" bson:"mobile"`
	Gender           string               `json:"gender" bson:"gender"`
	Dob              string               `json:"dob" bson:"dob"`
	PermanentAddress string               `json:"permanentAddress" bson:"permanentAddress"`
	PresentAddress   string               `json:"presentAddress" bson:"presentAddress"`
	Category         string               `json:"category" bson:"category"`
	FatherName       string               `json:"fatherName" bson:"fatherName"`
	FatherOccupation string               `json:"fatherOccupation" bson:"fatherOccupation"`
	MotherName       string               `json:"motherName" bson:"motherName"`
	MotherOccupation string               `json:"motherOccupation" bson:"motherOccupation"`
	MotherTongue     string               `json:"motherTongue" bson:"motherTongue"`
	EducationGap     string               `json:"educationGap" bson:"educationGap"`
	JeeRank          string               `json:"jeeRank" bson:"jeeRank"`
	Cgpa             float32              `json:"cgpa" bson:"cgpa"`
	ActiveBacklogs   int                  `json:"activeBacklogs" bson:"activeBacklogs"`
	TotalBacklogs    int                  `json:"totalBacklogs" bson:"totalBacklogs"`
	XBoard           string               `json:"xBoard" bson:"xBoard"`
	XYear            string               `json:"xYear" bson:"xYear"`
	XPercentage      float32              `json:"xPercentage" bson:"xPercentage"`
	XInstitute       string               `json:"xInstitute" bson:"xInstitute"`
	XiiBoard         string               `json:"xiiBoard" bson:"xiiBoard"`
	XiiYear          string               `json:"xiiYear" bson:"xiiYear"`
	XiiPercentage    float32              `json:"xiiPercentage" bson:"xiiPercentage"`
	XiiInstitute     string               `json:"xiiInstitute" bson:"xiiInstitute"`
	SemesterOne      float32              `json:"semesterOne" bson:"semesterOne"`
	SemesterTwo      float32              `json:"semesterTwo" bson:"semesterTwo"`
	SemesterThree    float32              `json:"semesterThree" bson:"semesterThree"`
	SemesterFour     float32              `json:"semesterFour" bson:"semesterFour"`
	SemesterFive     float32              `json:"semesterFive" bson:"semesterFive"`
	SemesterSix      float32              `json:"semesterSix" bson:"semesterSix"`
	UpdatedAt        primitive.DateTime   `json:"updatedAt" bson:"updatedAt"`
	CreatedAt        primitive.DateTime   `json:"createdAt" bson:"createdAt"`
}

type StudentPopulated struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Batch            int                `json:"batch" bson:"batch"`
	RollNo           int                `json:"rollNo" bson:"rollNo"`
	FirstName        string             `json:"firstName" bson:"firstName"`
	LastName         string             `json:"lastName" bson:"lastName"`
	Department       string             `json:"department" bson:"department"`
	Course           string             `json:"course" bson:"course"`
	Email            string             `json:"email" bson:"email"`
	PersonalEmail    string             `json:"personalEmail" bson:"personalEmail"`
	LinkedIn         string             `json:"linkedIn" bson:"linkedIn"`
	Github           string             `json:"github" bson:"github"`
	MicrosoftTeams   string             `json:"microsoftTeams" bson:"microsoftTeams"`
	Mobile           int64              `json:"mobile" bson:"mobile"`
	Gender           string             `json:"gender" bson:"gender"`
	Dob              string             `json:"dob" bson:"dob"`
	PermanentAddress string             `json:"permanentAddress" bson:"permanentAddress"`
	PresentAddress   string             `json:"presentAddress" bson:"presentAddress"`
	Category         string             `json:"category" bson:"category"`
	FatherName       string             `json:"fatherName" bson:"fatherName"`
	FatherOccupation string             `json:"fatherOccupation" bson:"fatherOccupation"`
	MotherName       string             `json:"motherName" bson:"motherName"`
	MotherOccupation string             `json:"motherOccupation" bson:"motherOccupation"`
	MotherTongue     string             `json:"motherTongue" bson:"motherTongue"`
	EducationGap     string             `json:"educationGap" bson:"educationGap"`
	JeeRank          string             `json:"jeeRank" bson:"jeeRank"`
	Cgpa             float32            `json:"cgpa" bson:"cgpa"`
	ActiveBacklogs   int                `json:"activeBacklogs" bson:"activeBacklogs"`
	TotalBacklogs    int                `json:"totalBacklogs" bson:"totalBacklogs"`
	XBoard           string             `json:"xBoard" bson:"xBoard"`
	XYear            string             `json:"xYear" bson:"xYear"`
	XPercentage      float32            `json:"xPercentage" bson:"xPercentage"`
	XInstitute       string             `json:"xInstitute" bson:"xInstitute"`
	XiiBoard         string             `json:"xiiBoard" bson:"xiiBoard"`
	XiiYear          string             `json:"xiiYear" bson:"xiiYear"`
	XiiPercentage    float32            `json:"xiiPercentage" bson:"xiiPercentage"`
	XiiInstitute     string             `json:"xiiInstitute" bson:"xiiInstitute"`
	SemesterOne      float32            `json:"semesterOne" bson:"semesterOne"`
	SemesterTwo      float32            `json:"semesterTwo" bson:"semesterTwo"`
	SemesterThree    float32            `json:"semesterThree" bson:"semesterThree"`
	SemesterFour     float32            `json:"semesterFour" bson:"semesterFour"`
	SemesterFive     float32            `json:"semesterFive" bson:"semesterFive"`
	SemesterSix      float32            `json:"semesterSix" bson:"semesterSix"`
	Groups           []group.Group      `json:"groups" bson:"groups"`
	UpdatedAt        primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	CreatedAt        primitive.DateTime `json:"createdAt" bson:"createdAt"`
	CompaniesAlloted []string           `json:"companiesAlloted" bson:"companiesAlloted"`
}
