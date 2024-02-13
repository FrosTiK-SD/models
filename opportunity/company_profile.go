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
	misc.Currency
	CompensationType CompensationType `bson:"compensationType" json:"compensationType"`
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
	APD *CompensationBreakup `bson:"APD" json:"APD"`
	CER *CompensationBreakup `bson:"CER" json:"CER"`
	CHE *CompensationBreakup `bson:"CHE" json:"CHE"`
	CIV *CompensationBreakup `bson:"CIV" json:"CIV"`
	CSE *CompensationBreakup `bson:"CSE" json:"CSE"`
	ECE *CompensationBreakup `bson:"ECE" json:"ECE"`
	EEE *CompensationBreakup `bson:"EEE" json:"EEE"`
	MEC *CompensationBreakup `bson:"MEC" json:"MEC"`
	MET *CompensationBreakup `bson:"MET" json:"MET"`
	MIN *CompensationBreakup `bson:"MIN" json:"MIN"`
	MST *CompensationBreakup `bson:"MST" json:"MST"`
	PHY *CompensationBreakup `bson:"PHY" json:"PHY"`
}

type IDDCompensation struct {
	APD *CompensationBreakup `bson:"APD" json:"APD"`
	BCE *CompensationBreakup `bson:"BCE" json:"BCE"`
	BME *CompensationBreakup `bson:"BME" json:"BME"`
	CER *CompensationBreakup `bson:"CER" json:"CER"`
	CHY *CompensationBreakup `bson:"CHY" json:"CHY"`
	CIV *CompensationBreakup `bson:"CIV" json:"CIV"`
	CSE *CompensationBreakup `bson:"CSE" json:"CSE"`
	ECE *CompensationBreakup `bson:"ECE" json:"ECE"`
	EEE *CompensationBreakup `bson:"EEE" json:"EEE"`
	MAT *CompensationBreakup `bson:"MAT" json:"MAT"`
	MEC *CompensationBreakup `bson:"MEC" json:"MEC"`
	MET *CompensationBreakup `bson:"MET" json:"MET"`
	MIN *CompensationBreakup `bson:"MIN" json:"MIN"`
	MST *CompensationBreakup `bson:"MST" json:"MST"`
	PHE *CompensationBreakup `bson:"PHE" json:"PHE"`
	PHY *CompensationBreakup `bson:"PHY" json:"PHY"`
}

type MTechCompensation struct {
	APD *CompensationBreakup `bson:"APD" json:"APD"`
	BCE *CompensationBreakup `bson:"BCE" json:"BCE"`
	BME *CompensationBreakup `bson:"BME" json:"BME"`
	CER *CompensationBreakup `bson:"CER" json:"CER"`
	CHE *CompensationBreakup `bson:"CHE" json:"CHE"`
	CHY *CompensationBreakup `bson:"CHY" json:"CHY"`
	CIV *CompensationBreakup `bson:"CIV" json:"CIV"`
	CSE *CompensationBreakup `bson:"CSE" json:"CSE"`
	DSE *CompensationBreakup `bson:"DSE" json:"DSE"`
	ECE *CompensationBreakup `bson:"ECE" json:"ECE"`
	EEE *CompensationBreakup `bson:"EEE" json:"EEE"`
	MAT *CompensationBreakup `bson:"MAT" json:"MAT"`
	MEC *CompensationBreakup `bson:"MEC" json:"MEC"`
	MET *CompensationBreakup `bson:"MET" json:"MET"`
	MIN *CompensationBreakup `bson:"MIN" json:"MIN"`
	MST *CompensationBreakup `bson:"MST" json:"MST"`
	PHE *CompensationBreakup `bson:"PHE" json:"PHE"`
	PHY *CompensationBreakup `bson:"PHY" json:"PHY"`
}

type PhDCompensation struct {
	APD *CompensationBreakup `bson:"APD" json:"APD"`
	BCE *CompensationBreakup `bson:"BCE" json:"BCE"`
	BME *CompensationBreakup `bson:"BME" json:"BME"`
	CER *CompensationBreakup `bson:"CER" json:"CER"`
	CHE *CompensationBreakup `bson:"CHE" json:"CHE"`
	CHY *CompensationBreakup `bson:"CHY" json:"CHY"`
	CIV *CompensationBreakup `bson:"CIV" json:"CIV"`
	CSE *CompensationBreakup `bson:"CSE" json:"CSE"`
	DSE *CompensationBreakup `bson:"DSE" json:"DSE"`
	ECE *CompensationBreakup `bson:"ECE" json:"ECE"`
	EEE *CompensationBreakup `bson:"EEE" json:"EEE"`
	MAT *CompensationBreakup `bson:"MAT" json:"MAT"`
	MEC *CompensationBreakup `bson:"MEC" json:"MEC"`
	MET *CompensationBreakup `bson:"MET" json:"MET"`
	MIN *CompensationBreakup `bson:"MIN" json:"MIN"`
	MST *CompensationBreakup `bson:"MST" json:"MST"`
	PHE *CompensationBreakup `bson:"PHE" json:"PHE"`
	PHY *CompensationBreakup `bson:"PHY" json:"PHY"`
}

type CourseCompensationDetails struct {
	BTech *BTechCompensation `bson:"BTECH" json:"BTECH"`
	IDD   *IDDCompensation   `bson:"IDD" json:"IDD"`
	MTech *MTechCompensation `bson:"MTECH" json:"MTECH"`
	PhD   *PhDCompensation   `bson:"PHD" json:"PHD"`
}

type CompanyProfile struct {
	Id                     primitive.ObjectID        `bson:"_id" json:"_id"`
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
