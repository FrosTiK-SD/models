package student

import (
	misc "github.com/FrosTiK-SD/models/misc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationCategory struct {
	Category string `json:"category" bson:"category"`
	IsPWD    bool   `json:"isPWD" bson:"isPWD"`
	IsEWS    bool   `json:"isEWS" bson:"isEWS"`
}

type RankDetails struct {
	Rank         int                  `bson:"rank" json:"rank"`
	RankCategory *ReservationCategory `bson:"rankCategory" json:"rankCategory"`
}

type SemesterSPI struct {
	One   *float64 `json:"one" bson:"one"`
	Two   *float64 `json:"two" bson:"two"`
	Three *float64 `json:"three" bson:"three"`
	Four  *float64 `json:"four" bson:"four"`
	Five  *float64 `json:"five" bson:"five"`
	Six   *float64 `json:"six" bson:"six"`
	Seven *float64 `json:"seven" bson:"seven"`
	Eight *float64 `json:"eight" bson:"eight"`
	Nine  *float64 `json:"nine" bson:"nine"`
	Ten   *float64 `json:"ten" bson:"ten"`
}

type SummerTermSPI struct {
	One   *float64 `json:"one" bson:"one"`
	Two   *float64 `json:"two" bson:"two"`
	Three *float64 `json:"three" bson:"three"`
	Four  *float64 `json:"four" bson:"four"`
	Five  *float64 `json:"five" bson:"five"`
}

type EducationDetails struct {
	Certification string  `bson:"certification" json:"certification"`
	Institute     string  `bson:"institute" json:"institute"`
	Year          int     `bson:"year" json:"year"`
	Score         float64 `bson:"score" json:"score"`
}

type DiplomaDetails struct {
	Board string `bson:"diploma_board" json:"diploma_board"`
	Institute string `bson:"diploma_from" json:"diploma_from"`
	Year int `bson:"diploma_year" json:"diploma_year"`
	Percentage float64 `bson:"diploma_percentage" json:"diploma_percentage"`
}

type Academics struct {
	JEERank        *RankDetails        `json:"jeeRank" bson:"jeeRank"`
	GATERank       *RankDetails        `json:"gateRank" bson:"gateRank"`
	XthClass       *EducationDetails   `json:"xClass" bson:"xClass"`
	XIIthClass     *EducationDetails   `json:"xiiClass" bson:"xiiClass"`
	UnderGraduate  *EducationDetails   `json:"underGraduate" bson:"underGraduate"`
	Honours        *string             `json:"honours" bson:"honours"`
	PostGraduate   *EducationDetails   `json:"postGraduate" bson:"postGraduate"`
	Diploma		   *DiplomaDetails	   `json:"diploma" bson:"diploma"`
	ThesisEndDate  *primitive.DateTime `json:"thesisEndDate" bson:"thesisEndDate"`
	EducationGap   *int                `json:"educationGap" bson:"educationGap"`
	SemesterSPI    SemesterSPI         `json:"semesterSPI" bson:"semesterSPI"`
	SummerTermSPI  SummerTermSPI       `json:"summerTermSPI" bson:"summerTermSPI"`
	CurrentCGPA    *float64            `json:"currentCGPA" bson:"currentCGPA"`
	ActiveBacklogs *int                `json:"activeBacklogs" bson:"activeBacklogs"`
	TotalBacklogs  *int                `json:"totalBacklogs" bson:"totalBacklogs"`

	Verification misc.Verification `json:"verification" bson:"verification"`
}

type SocialProfile struct {
	URL          string            `bson:"url" json:"url"`
	Username     string            `bson:"username" json:"username"`
	Verification misc.Verification `bson:"verification" json:"verification"`
}

type SocialProfiles struct {
	LinkedIn       *SocialProfile `json:"linkedIn" bson:"linkedIn"`
	Github         *SocialProfile `json:"github" bson:"github"`
	MicrosoftTeams *SocialProfile `json:"microsoftTeams" bson:"microsoftTeams"`
	Skype          *SocialProfile `json:"skype" bson:"skype"`
	GoogleScholar  *SocialProfile `json:"googleScholar" bson:"googleScholar"`

	Codeforces *SocialProfile `json:"codeforces" bson:"codeforces"`
	CodeChef   *SocialProfile `json:"codechef" bson:"codechef"`
	LeetCode   *SocialProfile `json:"leetcode" bson:"leetcode"`
	Kaggle     *SocialProfile `json:"kaggle" bson:"kaggle"`
}

type Batch struct {
	StartYear int `json:"startYear" bson:"startYear"`
	EndYear   int `json:"endYear" bson:"endYear"`
}

type ParentsDetails struct {
	FatherName       string `json:"fatherName" bson:"fatherName"`
	FatherOccupation string `json:"fatherOccupation" bson:"fatherOccupation"`
	MotherName       string `json:"motherName" bson:"motherName"`
	MotherOccupation string `json:"motherOccupation" bson:"motherOccupation"`
}

type Extras struct {
	VideoResume  *string           `bson:"videoResume" json:"videoResume"`
	Verification misc.Verification `bson:"verification" json:"verification"`
}

type WorkExperience struct {
	StartDate    primitive.DateTime `bson:"startDate" json:"startDate"`
	EndDate      primitive.DateTime `bson:"endDate" json:"endDate"`
	Organisation string             `bson:"organisation" json:"organisation"`
	Location     string             `bson:"location" json:"location"`
	Position     string             `bson:"position" json:"position"`
	Details      string             `bson:"details" json:"details"`
	Verification misc.Verification  `bson:"verification" json:"verification"`
}

type UserError string

const (
	MIGRATION UserError = "migration"
)

type DataErrors struct {
	User map[string]UserError `json:"user" bson:"user"`
}
