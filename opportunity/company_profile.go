package opportunity

import (
	"github.com/FrosTiK-SD/iitbhu-tpc-models-golang/misc"
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

type CompanyProfile struct {
	ID                     primitive.ObjectID     `bson:"_id" json:"_id,omitempty"`
	Company                primitive.ObjectID     `bson:"company" json:"company,omitempty"`
	Criterias              []primitive.ObjectID   `bson:"criterias" json:"criterias,omitempty"`
	Attachments            []misc.Attachment      `bson:"attachments" json:"attachments,omitempty"`
	Slot                   primitive.ObjectID     `bson:"slot" json:"slot,omitempty"`
	DetailsRequestedSchema map[string]interface{} `bson:"detailsRequestedSchema" json:"detailsRequestedSchema,omitempty"`

	ApplicationStartTime primitive.DateTime `bson:"applicationStartTime" json:"applicationStartTime,omitempty"`
	ApplicationEndTime   primitive.DateTime `bson:"applicationEndTime" json:"applicationEndTime,omitempty"`
	ListingStatus        ListingStatus      `bson:"listingStatus" json:"listingStatus,omitempty"`
	ListingType          ListingType        `bson:"listingType" json:"listingType,omitempty"`

	// metadata
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt" json:"updatedAt,omitempty"`
}
