package student

import (
	constant "github.com/FrosTiK-SD/models/constant"
	misc "github.com/FrosTiK-SD/models/misc"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	Id               primitive.ObjectID   `json:"_id" bson:"_id"`
	Groups           []primitive.ObjectID `json:"groups" bson:"groups"`
	CompaniesAlloted []string             `json:"companiesAlloted" bson:"companiesAlloted"`

	Batch          *Batch           `json:"batch" bson:"batch"`
	RollNo         int              `json:"rollNo" bson:"rollNo"`
	InstituteEmail string           `json:"email" bson:"email"`
	Department     string           `json:"department" bson:"department"`
	Course         *constant.Course `json:"course" bson:"course"`
	Specialisation *string          `json:"specialisation" bson:"specialisation"`

	FirstName  string  `json:"firstName" bson:"firstName"`
	MiddleName *string `json:"middleName" bson:"middleName"`
	LastName   *string `json:"lastName" bson:"lastName"`

	ProfilePicture   *misc.Attachment     `json:"profilePicture" bson:"profilePicture"`
	Gender           *constant.Gender     `json:"gender" bson:"gender"`
	DOB              *primitive.DateTime  `json:"dob" bson:"dob"`
	PermanentAddress string               `json:"permanentAddress" bson:"permanentAddress"`
	PresentAddress   string               `json:"presentAddress" bson:"presentAddress"`
	PersonalEmail    string               `json:"personalEmail" bson:"personalEmail"`
	Mobile           string               `json:"mobile" bson:"mobile"`
	Category         *ReservationCategory `json:"category" bson:"category"`
	MotherTongue     string               `json:"motherTongue" bson:"motherTongue"`
	ParentsDetails   ParentsDetails       `json:"parentsDetails" bson:"parentsDetails"`

	Academics      Academics        `json:"academics" bson:"academics"`
	WorkExperience []WorkExperience `json:"workExperience" bson:"workExperience"`
	SocialProfiles SocialProfiles   `json:"socialProfiles" bson:"socialProfiles"`
	Extras         Extras           `json:"extras" bson:"extras"`

	// metadata
	StructVersion int                    `json:"version" bson:"version"`
	UpdatedAt     primitive.DateTime     `json:"updatedAt" bson:"updatedAt"`
	CreatedAt     primitive.DateTime     `json:"createdAt" bson:"createdAt"`
	RawKeyStore   map[string]interface{} `json:"raw_key_store" bson:"raw_key_store"`
	DataErrors    DataErrors             `json:"dataErrors" bson:"dataErrors"`
}
