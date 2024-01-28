package student

import (
	constant "github.com/FrosTiK-SD/models/constant"
	misc "github.com/FrosTiK-SD/models/misc"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationCategory struct {
	Category string `json:"category,omitempty" bson:"category"`
	IsPWD    bool   `json:"isPWD,omitempty" bson:"isPWD"`
	IsEWS    bool   `json:"isEWS,omitempty" bson:"isEWS"`
}

type RankDetails struct {
	Rank         int                 `bson:"rank" json:"rank,omitempty"`
	RankCategory ReservationCategory `bson:"rankCategory" json:"rankCategory,omitempty"`
}

type SemesterSPI struct {
	One   float32 `json:"One,omitempty" bson:"One"`
	Two   float32 `json:"Two,omitempty" bson:"Two"`
	Three float32 `json:"Three,omitempty" bson:"Three"`
	Four  float32 `json:"Four,omitempty" bson:"Four"`
	Five  float32 `json:"Five,omitempty" bson:"Five"`
	Six   float32 `json:"Six,omitempty" bson:"Six"`
	Seven float32 `json:"Seven,omitempty" bson:"Seven"`
	Eight float32 `json:"Eight,omitempty" bson:"Eight"`
	Nine  float32 `json:"Nine,omitempty" bson:"Nine"`
	Ten   float32 `json:"Ten,omitempty" bson:"Ten"`
}

type SummerTermSPI struct {
	One   float32 `json:"One,omitempty" bson:"One"`
	Two   float32 `json:"Two,omitempty" bson:"Two"`
	Three float32 `json:"Three,omitempty" bson:"Three"`
	Four  float32 `json:"Four,omitempty" bson:"Four"`
	Five  float32 `json:"Five,omitempty" bson:"Five"`
}

type EducationDetails struct {
	Certification string  `bson:"certification" json:"certification,omitempty"`
	Institute     string  `bson:"institute" json:"institute,omitempty"`
	Year          int     `bson:"year" json:"year,omitempty"`
	Score         float32 `bson:"score" json:"score,omitempty"`
}

type Academics struct {
	JEERank           RankDetails        `json:"jeeRank" bson:"jeeRank"`
	GATERank          RankDetails        `json:"gateRank" bson:"gateRank"`
	XthClass          EducationDetails   `json:"xClass" bson:"xClass"`
	XIIthClass        EducationDetails   `json:"xiiClass" bson:"xiiClass"`
	UnderGraduate     EducationDetails   `json:"underGraduate" bson:"underGraduate"`
	Honours           string             `json:"honours" bson:"honours"`
	PostGraduate      EducationDetails   `json:"postGraduate" bson:"postGraduate"`
	ThesisEndDate     primitive.DateTime `json:"thesisEndDate" bson:"thesisEndDate"`
	EducationGap      int                `json:"educationGap" bson:"educationGap"`
	SemesterDetails   SemesterSPI        `json:"semesterSPI" bson:"semesterSPI"`
	SummerTermDetails SummerTermSPI      `json:"summerTermSPI" bson:"summerTermSPI"`
	CurrentCGPA       float32            `json:"currentCGPA,omitempty" bson:"currentCGPA"`
	ActiveBacklogs    int                `json:"activeBacklogs" bson:"activeBacklogs"`
	TotalBacklogs     int                `json:"totalBacklogs" bson:"totalBacklogs"`

	Verification misc.Verification `json:"verification" bson:"verification"`
}

type SocialProfile struct {
	URL          string            `bson:"url" json:"url,omitempty"`
	Username     string            `bson:"username" json:"username,omitempty"`
	Verification misc.Verification `bson:"verification" json:"verification,omitempty"`
}

type SocialProfiles struct {
	LinkedIn       SocialProfile `json:"linkedIn" bson:"linkedIn"`
	Github         SocialProfile `json:"github" bson:"github"`
	MicrosoftTeams SocialProfile `json:"microsoftTeams" bson:"microsoftTeams"`
	Skype          SocialProfile `json:"skype" bson:"skype"`
	GoogleScholar  SocialProfile `json:"googleScholar" bson:"googleScholar"`

	Codeforces SocialProfile `json:"codeforces" bson:"codeforces"`
	CodeChef   SocialProfile `json:"codechef" bson:"codechef"`
	LeetCode   SocialProfile `json:"leetcode" bson:"leetcode"`
	Kaggle     SocialProfile `json:"kaggle" bson:"kaggle"`
}

type Batch struct {
	StartYear int `json:"startYear,omitempty" bson:"startYear"`
	EndYear   int `json:"endYear,omitempty" bson:"endYear"`
}

type ParentsDetails struct {
	FatherName       string `json:"fatherName" bson:"fatherName"`
	FatherOccupation string `json:"fatherOccupation" bson:"fatherOccupation"`
	MotherName       string `json:"motherName" bson:"motherName"`
	MotherOccupation string `json:"motherOccupation" bson:"motherOccupation"`
}

type Extras struct {
	VideoResume  string            `bson:"videoResumes" json:"videoResume,omitempty"`
	Verification misc.Verification `bson:"verification" json:"verification,omitempty"`
}

type WorkExperience struct {
	StartDate    primitive.DateTime `bson:"startDate" json:"startDate,omitempty"`
	EndDate      primitive.DateTime `bson:"endDate" json:"endDate,omitempty"`
	Organisation string             `bson:"organisation" json:"organisation,omitempty"`
	Location     string             `bson:"location" json:"location,omitempty"`
	Position     string             `bson:"position" json:"position"`
	Details      string             `bson:"details" json:"details,omitempty"`
	Verification misc.Verification  `bson:"verification" json:"verification,omitempty"`
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

	FirstName  string `json:"firstName" bson:"firstName"`
	MiddleName string `json:"middleName" bson:"middleName"`
	LastName   string `json:"lastName" bson:"lastName"`

	ProfilePicture   misc.Attachment     `json:"profilePicture" bson:"profilePicture"`
	Gender           constant.Gender     `json:"gender" bson:"gender"`
	DOB              primitive.DateTime  `json:"dob" bson:"dob"`
	PermanentAddress string              `json:"permanentAddress" bson:"permanentAddress"`
	PresentAddress   string              `json:"presentAddress" bson:"presentAddress"`
	PersonalEmail    string              `json:"personalEmail" bson:"personalEmail"`
	Mobile           string              `json:"mobile" bson:"mobile"`
	Category         ReservationCategory `json:"category" bson:"category"`
	MotherTongue     string              `json:"motherTongue" bson:"motherTongue"`
	ParentsDetails   ParentsDetails      `json:"parentsDetails" bson:"parentsDetails"`

	Academics      Academics        `json:"academics" bson:"academics"`
	WorkExperience []WorkExperience `json:"workExperience" bson:"workExperience"`
	SocialProfiles SocialProfiles   `json:"socialProfiles" bson:"socialProfiles"`

	// metadata
	UpdatedAt   primitive.DateTime     `json:"updatedAt" bson:"updatedAt"`
	CreatedAt   primitive.DateTime     `json:"createdAt" bson:"createdAt"`
	RawKeyStore map[string]interface{} `json:"raw_key_store" bson:"raw_key_store"`
}
