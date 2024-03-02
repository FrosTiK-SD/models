package opportunity

import (
	"github.com/FrosTiK-SD/models/misc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListingStatus string
type ListingType string
type CompensationType string

const (
	LISTED             ListingStatus = "LISTED"
	UNLISTED           ListingStatus = "UNLISTED"
	UNDER_BRANCH_ISSUE ListingStatus = "UNDER_BRANCH_ISSUE"
	CLOSED             ListingStatus = "CLOSED"
)

const (
	PLACEMENT ListingType = "PLACEMENT"
	INTERN    ListingType = "INTERN"
)

const (
	PER_YEAR  CompensationType = "PER_YEAR"
	PER_MONTH CompensationType = "PER_MONTH"
)

type Deadlines struct {
	PPTDate       *primitive.DateTime `bson:"pptDate" json:"pptDate"`
	ExamDate      *primitive.DateTime `bson:"examDate" json:"examDate"`
	InterviewDate *primitive.DateTime `bson:"interviewDate" json:"interviewDate"`
}

type Compensation struct {
	Currency         misc.CurrencyType `json:"currency" bson:"currency"`
	Amount           float64           `json:"amount" bson:"amount"`
	CompensationType CompensationType  `bson:"compensationType" json:"compensationType"`
}

type CompensationRange struct {
	Min *Compensation `bson:"min" json:"min"`
	Max *Compensation `bson:"max" json:"max"`
}

type CompensationBreakup struct {
	TotalCTC *CompensationRange `bson:"totalCTC" json:"totalCTC"`
	Fixed    *CompensationRange `bson:"fixed" json:"fixed"`
	Extras   *string            `bson:"extras" json:"extras"`
}

type BTechCompensation struct {
	APD *CompensationBreakup `bson:"apd" json:"apd"`
	CER *CompensationBreakup `bson:"cer" json:"cer"`
	CHE *CompensationBreakup `bson:"che" json:"che"`
	CIV *CompensationBreakup `bson:"civ" json:"civ"`
	CSE *CompensationBreakup `bson:"cse" json:"cse"`
	ECE *CompensationBreakup `bson:"ece" json:"ece"`
	EEE *CompensationBreakup `bson:"eee" json:"eee"`
	MEC *CompensationBreakup `bson:"mec" json:"mec"`
	MET *CompensationBreakup `bson:"met" json:"met"`
	MIN *CompensationBreakup `bson:"min" json:"min"`
	MST *CompensationBreakup `bson:"mst" json:"mst"`
	PHY *CompensationBreakup `bson:"phy" json:"phy"`
}

type IDDCompensation struct {
	APD *CompensationBreakup `bson:"apd" json:"apd"`
	BCE *CompensationBreakup `bson:"bce" json:"bce"`
	BME *CompensationBreakup `bson:"bme" json:"bme"`
	CER *CompensationBreakup `bson:"cer" json:"cer"`
	CHY *CompensationBreakup `bson:"chy" json:"chy"`
	CIV *CompensationBreakup `bson:"civ" json:"civ"`
	CSE *CompensationBreakup `bson:"cse" json:"cse"`
	ECE *CompensationBreakup `bson:"ece" json:"ece"`
	EEE *CompensationBreakup `bson:"eee" json:"eee"`
	MAT *CompensationBreakup `bson:"mat" json:"mat"`
	MEC *CompensationBreakup `bson:"mec" json:"mec"`
	MET *CompensationBreakup `bson:"met" json:"met"`
	MIN *CompensationBreakup `bson:"min" json:"min"`
	MST *CompensationBreakup `bson:"mst" json:"mst"`
	PHE *CompensationBreakup `bson:"phe" json:"phe"`
	PHY *CompensationBreakup `bson:"phy" json:"phy"`
}

type MTechCompensation struct {
	APD *CompensationBreakup `bson:"apd" json:"apd"`
	BCE *CompensationBreakup `bson:"bce" json:"bce"`
	BME *CompensationBreakup `bson:"bme" json:"bme"`
	CER *CompensationBreakup `bson:"cer" json:"cer"`
	CHE *CompensationBreakup `bson:"che" json:"che"`
	CHY *CompensationBreakup `bson:"chy" json:"chy"`
	CIV *CompensationBreakup `bson:"civ" json:"civ"`
	CSE *CompensationBreakup `bson:"cse" json:"cse"`
	DSE *CompensationBreakup `bson:"dse" json:"dse"`
	ECE *CompensationBreakup `bson:"ece" json:"ece"`
	EEE *CompensationBreakup `bson:"eee" json:"eee"`
	MAT *CompensationBreakup `bson:"mat" json:"mat"`
	MEC *CompensationBreakup `bson:"mec" json:"mec"`
	MET *CompensationBreakup `bson:"met" json:"met"`
	MIN *CompensationBreakup `bson:"min" json:"min"`
	MST *CompensationBreakup `bson:"mst" json:"mst"`
	PHE *CompensationBreakup `bson:"phe" json:"phe"`
	PHY *CompensationBreakup `bson:"phy" json:"phy"`
}

type PhDCompensation struct {
	APD *CompensationBreakup `bson:"apd" json:"apd"`
	BCE *CompensationBreakup `bson:"bce" json:"bce"`
	BME *CompensationBreakup `bson:"bme" json:"bme"`
	CER *CompensationBreakup `bson:"cer" json:"cer"`
	CHE *CompensationBreakup `bson:"che" json:"che"`
	CHY *CompensationBreakup `bson:"chy" json:"chy"`
	CIV *CompensationBreakup `bson:"civ" json:"civ"`
	CSE *CompensationBreakup `bson:"cse" json:"cse"`
	DSE *CompensationBreakup `bson:"dse" json:"dse"`
	ECE *CompensationBreakup `bson:"ece" json:"ece"`
	EEE *CompensationBreakup `bson:"eee" json:"eee"`
	MAT *CompensationBreakup `bson:"mat" json:"mat"`
	MEC *CompensationBreakup `bson:"mec" json:"mec"`
	MET *CompensationBreakup `bson:"met" json:"met"`
	MIN *CompensationBreakup `bson:"min" json:"min"`
	MST *CompensationBreakup `bson:"mst" json:"mst"`
	PHE *CompensationBreakup `bson:"phe" json:"phe"`
	PHY *CompensationBreakup `bson:"phy" json:"phy"`
}

type CourseCompensationDetails struct {
	BTech *BTechCompensation `bson:"BTech" json:"BTech"`
	IDD   *IDDCompensation   `bson:"IDD" json:"IDD"`
	MTech *MTechCompensation `bson:"MTech" json:"MTech"`
	PhD   *PhDCompensation   `bson:"PhD" json:"PhD"`
}

type CompanyProfile struct {
	Id                     primitive.ObjectID        `bson:"_id" json:"_id"`
	Opportunity            *primitive.ObjectID       `bson:"opportunity" json:"opportunity"`
	Company                *primitive.ObjectID       `bson:"company" json:"company"`
	Role                   *string                   `bson:"role" json:"role"`
	Batch                  *string                   `bson:"batch" json:"batch"`
	Deadlines              Deadlines                 `bson:"deadlines" json:"deadlines"`
	CompensationDetails    CourseCompensationDetails `bson:"compensationDetails" json:"compensationDetails"`
	JD                     *string                   `bson:"jd" json:"jd"`
	Criterias              *[]primitive.ObjectID     `bson:"criterias" json:"criterias"`
	Spocs                  *[]primitive.ObjectID     `bson:"spocs" json:"spocs"`
	Attachments            []misc.Attachment         `bson:"attachments" json:"attachments,omitempty"`
	ApplicationStartTime   primitive.DateTime        `bson:"applicationStartTime" json:"applicationStartTime,omitempty"`
	ApplicationEndTime     primitive.DateTime        `bson:"applicationEndTime" json:"applicationEndTime,omitempty"`
	Slot                   *primitive.ObjectID       `bson:"slot" json:"slot"`
	DetailsRequestedSchema *map[string]interface{}   `bson:"detailsRequestedSchema" json:"detailsRequestedSchema"`
	ListingStatus          ListingStatus             `bson:"listingStatus" json:"listingStatus,omitempty"`
	ListingType            ListingType               `bson:"listingType" json:"listingType,omitempty"`
	Activities             []misc.Activity           `bson:"activities" json:"activities"`
	CreatedAt              primitive.DateTime        `bson:"createdAt" json:"createdAt"`
	UpdatedAt              primitive.DateTime        `bson:"updatedAt" json:"updatedAt"`
}
