package opportunity

import (
	"github.com/FrosTiK-SD/models/misc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListingStatus string
type ListingType string

const (
	LISTED             ListingStatus = "listed"
	UNLISTED           ListingStatus = "unlisted"
	UNDER_BRANCH_ISSUE ListingStatus = "under_branch_issue"
	CLOSED             ListingStatus = "closed"
)

const (
	PLACEMENT ListingType = "placement"
	INTERN    ListingType = "intern"
)

type Deadlines struct {
	PPTDate       *primitive.DateTime `bson:"pptDate" json:"pptDate"`
	ExamDate      *primitive.DateTime `bson:"examDate" json:"examDate"`
	InterviewDate *primitive.DateTime `bson:"interviewDate" json:"interviewDate"`
}

type CompensationBreakup struct {
	TotalCTC *misc.Currency `bson:"totalCTC" json:"totalCTC"`
	Fixed    *misc.Currency `bson:"fixed" json:"fixed"`
	Extras   *string        `bson:"extras" json:"extras"`
}

type CompensationDetails struct {
	BTech *CompensationBreakup `bson:"btech" json:"btech"`
	IDD   *CompensationBreakup `bson:"idd" json:"idd"`
	MTech *CompensationBreakup `bson:"mtech" json:"mtech"`
	PhD   *CompensationBreakup `bson:"phd" json:"phd"`
}

type CompanyProfile struct {
	Id                     primitive.ObjectID      `bson:"_id" json:"_id"`
	Company                *primitive.ObjectID     `bson:"company" json:"company"`
	Role                   *string                 `bson:"role" json:"role"`
	Batch                  *string                 `bson:"batch" json:"batch"`
	Deadlines              Deadlines               `bson:"deadlines" json:"deadlines"`
	CompensationDetails    CompensationDetails     `bson:"compensationDetails" json:"compensationDetails"`
	JD                     *string                 `bson:"jd" json:"jd"`
	Criterias              *[]primitive.ObjectID   `bson:"criterias" json:"criterias"`
	Spocs                  *[]primitive.ObjectID   `bson:"spocs" json:"spocs"`
	Attachments            []misc.Attachment       `bson:"attachments" json:"attachments,omitempty"`
	ApplicationStartTime   primitive.DateTime      `bson:"applicationStartTime" json:"applicationStartTime,omitempty"`
	ApplicationEndTime     primitive.DateTime      `bson:"applicationEndTime" json:"applicationEndTime,omitempty"`
	Slot                   *primitive.ObjectID     `bson:"slot" json:"slot"`
	DetailsRequestedSchema *map[string]interface{} `bson:"detailsRequestedSchema" json:"detailsRequestedSchema"`
	ListingStatus          ListingStatus           `bson:"listingStatus" json:"listingStatus,omitempty"`
	ListingType            ListingType             `bson:"listingType" json:"listingType,omitempty"`
	Activities             []misc.Activity         `bson:"activities" json:"activities"`
	CreatedAt              primitive.DateTime      `bson:"createdAt" json:"createdAt"`
	UpdatedAt              primitive.DateTime      `bson:"updatedAt" json:"updatedAt"`
}
