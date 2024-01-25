package student

import (
	constant "github.com/FrosTiK-SD/models/constant"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SemesterCGPA struct {
	One   float32 `json:"one,omitempty" bson:"one"`
	Two   float32 `json:"two,omitempty" bson:"two"`
	Three float32 `json:"three,omitempty" bson:"three"`
	Four  float32 `json:"four,omitempty" bson:"four"`
	Five  float32 `json:"five,omitempty" bson:"five"`
	Six   float32 `json:"six,omitempty" bson:"six"`
	Seven float32 `json:"seven,omitempty" bson:"seven"`
	Eight float32 `json:"eight,omitempty" bson:"eight"`
}

type Batch struct {
	StartYear int `json:"startYear,omitempty" bson:"startYear"`
	EndYear   int `json:"endYear,omitempty" bson:"endYear"`
}

type Student struct {
	ID               primitive.ObjectID           `json:"_id" bson:"_id"`
	Groups           []primitive.ObjectID         `json:"groups" bson:"groups"`
	CompaniesAlloted []string                     `json:"companiesAlloted" bson:"companiesAlloted"`
	Batch            Batch                        `json:"batch" bson:"batch"`
	RollNo           int                          `json:"rollNo" bson:"rollNo"`
	FirstName        string                       `json:"firstName" bson:"firstName"`
	LastName         string                       `json:"lastName" bson:"lastName"`
	Department       string                       `json:"department" bson:"department"`
	Course           constant.Course              `json:"course" bson:"course"`
	Email            string                       `json:"email" bson:"email"`
	PersonalEmail    string                       `json:"personalEmail" bson:"personalEmail"`
	LinkedIn         string                       `json:"linkedIn" bson:"linkedIn"`
	Github           string                       `json:"github" bson:"github"`
	MicrosoftTeams   string                       `json:"microsoftTeams" bson:"microsoftTeams"`
	Mobile           string                       `json:"mobile" bson:"mobile"`
	Gender           constant.Gender              `json:"gender" bson:"gender"`
	DOB              string                       `json:"dob" bson:"dob"`
	PermanentAddress string                       `json:"permanentAddress" bson:"permanentAddress"`
	PresentAddress   string                       `json:"presentAddress" bson:"presentAddress"`
	Category         constant.ReservationCategory `json:"category" bson:"category"`
	FatherName       string                       `json:"fatherName" bson:"fatherName"`
	FatherOccupation string                       `json:"fatherOccupation" bson:"fatherOccupation"`
	MotherName       string                       `json:"motherName" bson:"motherName"`
	MotherOccupation string                       `json:"motherOccupation" bson:"motherOccupation"`
	MotherTongue     string                       `json:"motherTongue" bson:"motherTongue"`
	EducationGap     int                          `json:"educationGap" bson:"educationGap"`
	JEERank          int                          `json:"jeeRank" bson:"jeeRank"`
	CGPA             float32                      `json:"cgpa" bson:"cgpa"`
	ActiveBacklogs   int                          `json:"activeBacklogs" bson:"activeBacklogs"`
	TotalBacklogs    int                          `json:"totalBacklogs" bson:"totalBacklogs"`
	XBoard           string                       `json:"xBoard" bson:"xBoard"`
	XYear            int                          `json:"xYear" bson:"xYear"`
	XPercentage      float32                      `json:"xPercentage" bson:"xPercentage"`
	XInstitute       string                       `json:"xInstitute" bson:"xInstitute"`
	XiiBoard         string                       `json:"xiiBoard" bson:"xiiBoard"`
	XiiYear          int                          `json:"xiiYear" bson:"xiiYear"`
	XiiPercentage    float32                      `json:"xiiPercentage" bson:"xiiPercentage"`
	XiiInstitute     string                       `json:"xiiInstitute" bson:"xiiInstitute"`
	SemesterCGPA     SemesterCGPA                 `json:"semesterCGPA" bson:"semesterCGPA"`

	// metadata
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
}
