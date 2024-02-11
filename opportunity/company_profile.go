package opportunity

import (
	"github.com/FrosTiK-SD/models/constant"
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
	PER_YEAR  ListingType = "PER_YEAR"
	PER_MONTH ListingType = "PER_MONTH"
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

type BranchCompensationDetails map[constant.Branch]*CompensationBreakup

type CourseCompensationDetails struct {
	BTech *BranchCompensationDetails `bson:"btech" json:"btech"`
	IDD   *BranchCompensationDetails `bson:"idd" json:"idd"`
	MTech *BranchCompensationDetails `bson:"mtech" json:"mtech"`
	PhD   *BranchCompensationDetails `bson:"phd" json:"phd"`
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
