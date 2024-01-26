package student

import (
	constant "github.com/FrosTiK-SD/models/constant"
	misc "github.com/FrosTiK-SD/models/misc"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationCategory struct {
	Category string `json:"category,omitempty" bson:"category"`
	IsPWD    bool   `json:"isPwd,omitempty" bson:"isPwD"`
	IsEWS    bool   `json:"isEWS,omitempty" bson:"isEWS"`
}

type RankDetails struct {
	Rank         int                 `bson:"rank" json:"rank,omitempty"`
	RankCategory ReservationCategory `bson:"rankCategory" json:"rankCategory,omitempty"`
}

type Academics struct {
	JEERank  RankDetails `json:"jeeRank" bson:"jeeRank"`
	GATERank RankDetails `json:"gateRank" bson:"gateRank"`

	XBoard      string  `json:"xBoard" bson:"xBoard"`
	XYear       int     `json:"xYear" bson:"xYear"`
	XPercentage float32 `json:"xPercentage" bson:"xPercentage"`
	XInstitute  string  `json:"xInstitute" bson:"xInstitute"`

	XiiBoard      string  `json:"xiiBoard" bson:"xiiBoard"`
	XiiYear       int     `json:"xiiYear" bson:"xiiYear"`
	XiiPercentage float32 `json:"xiiPercentage" bson:"xiiPercentage"`
	XiiInstitute  string  `json:"xiiInstitute" bson:"xiiInstitute"`

	EducationGap int `json:"educationGap" bson:"educationGap"`

	SemesterOne   float32 `json:"semesterOne,omitempty" bson:"semesterOne"`
	SemesterTwo   float32 `json:"semesterTwo,omitempty" bson:"semesterTwo"`
	SemesterThree float32 `json:"semesterThree,omitempty" bson:"semesterThree"`
	SemesterFour  float32 `json:"semesterFour,omitempty" bson:"semesterFour"`
	SemesterFive  float32 `json:"semesterFive,omitempty" bson:"semesterFive"`
	SemesterSix   float32 `json:"semesterSix,omitempty" bson:"semesterSix"`
	SemesterSeven float32 `json:"semesterSeven,omitempty" bson:"semesterSeven"`
	SemesterEight float32 `json:"semesterEight,omitempty" bson:"semesterEight"`
	SemesterNine  float32 `json:"semesterNine,omitempty" bson:"semesterNine"`
	SemesterTen   float32 `json:"semesterTen,omitempty" bson:"semesterTen"`

	SummerTermOne   float32 `json:"summerTermOne,omitempty" bson:"summerTermOne"`
	SummerTermTwo   float32 `json:"summerTermTwo,omitempty" bson:"summerTermTwo"`
	SummerTermThree float32 `json:"summerTermThree,omitempty" bson:"summerTermThree"`
	SummerTermFour  float32 `json:"summerTermFour,omitempty" bson:"summerTermFour"`
	SummerTermFive  float32 `json:"summerTermFive,omitempty" bson:"summerTermFive"`

	CurrentCGPA    float32 `json:"currentCGPA,omitempty" bson:"currentCGPA"`
	ActiveBacklogs int     `json:"activeBacklogs" bson:"activeBacklogs"`
	TotalBacklogs  int     `json:"totalBacklogs" bson:"totalBacklogs"`

	Verification misc.Verification `json:"verification" bson:"verification"`
}

type SocialProfiles struct {
	PersonalEmail  string `json:"personalEmail" bson:"personalEmail"`
	Mobile         string `json:"mobile" bson:"mobile"`
	LinkedIn       string `json:"linkedIn" bson:"linkedIn"`
	Github         string `json:"github" bson:"github"`
	MicrosoftTeams string `json:"microsoftTeams" bson:"microsoftTeams"`
	GoogleScholar  string `json:"googleScholar" bson:"googleScholar"`

	Codeforces string `json:"codeforces" bson:"codeforces"`
	CodeChef   string `json:"codechef" bson:"codechef"`
	LeetCode   string `json:"leetcode" bson:"leetcode"`
	Kaggle     string `json:"kaggle" bson:"kaggle"`

	Verification misc.Verification `json:"verification" bson:"verification"`
}

type Batch struct {
	StartYear int `json:"startYear,omitempty" bson:"startYear"`
	EndYear   int `json:"endYear,omitempty" bson:"endYear"`
}

type Student struct {
	ID               primitive.ObjectID   `json:"_id" bson:"_id"`
	Groups           []primitive.ObjectID `json:"groups" bson:"groups"`
	CompaniesAlloted []string             `json:"companiesAlloted" bson:"companiesAlloted"`

	Batch          Batch           `json:"batch" bson:"batch"`
	RollNo         int             `json:"rollNo" bson:"rollNo"`
	InstituteEmail string          `json:"instituteEmail" bson:"instituteEmail"`
	Department     string          `json:"department" bson:"department"`
	Course         constant.Course `json:"course" bson:"course"`
	Specialisation string          `json:"specialisation" bson:"specialisation"`

	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`

	Gender           constant.Gender     `json:"gender" bson:"gender"`
	DOB              string              `json:"dob" bson:"dob"`
	PermanentAddress string              `json:"permanentAddress" bson:"permanentAddress"`
	PresentAddress   string              `json:"presentAddress" bson:"presentAddress"`
	Category         ReservationCategory `json:"category" bson:"category"`
	FatherName       string              `json:"fatherName" bson:"fatherName"`
	FatherOccupation string              `json:"fatherOccupation" bson:"fatherOccupation"`
	MotherName       string              `json:"motherName" bson:"motherName"`
	MotherOccupation string              `json:"motherOccupation" bson:"motherOccupation"`
	MotherTongue     string              `json:"motherTongue" bson:"motherTongue"`

	Academics      Academics      `json:"academics" bson:"academics"`
	SocialProfiles SocialProfiles `json:"socialProfiles" bson:"socialProfiles"`

	// metadata
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
}
