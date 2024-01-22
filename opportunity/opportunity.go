package opportunity

import (
	"github.com/FrosTiK-SD/iitbhu-tpc-models-golang/company"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Listing string
type OpportunityType string

const (
	LISTED             Listing = "listed"
	UNLISTED           Listing = "unlisted"
	UNDER_BRANCH_ISSUE Listing = "under_branch_issue"
	CLOSED             Listing = "closed"
)

const (
	PLACEMENT OpportunityType = "placement"
	INTERN    OpportunityType = "intern"
)

type Opportunity struct {
	ID                        primitive.ObjectID          `bson:"id" json:"id,omitempty"`
	Company                   primitive.ObjectID          `bson:"company" json:"company,omitempty"`
	Profile                   string                      `bson:"profile" json:"profile,omitempty"`
	Criteria                  company.EligibilityCriteria `bson:"criteria" json:"criteria,omitempty"`
	DetailsRequestedGenerator map[string]interface{}      `bson:"details_requested_generator" json:"details_requested_generator,omitempty"` // formbuilder
	OpportunityStartTime      primitive.DateTime          `bson:"opportunity_start_time" json:"opportunity_start_time,omitempty"`
	OpportunityEndTime        primitive.DateTime          `bson:"opportunity_end_time" json:"opportunity_end_time,omitempty"`
	OpportunityStatus         Listing                     `bson:"listing" json:"listing,omitempty"`
	OpportunityType           OpportunityType             `bson:"opportunity_type" json:"opportunity_type,omitempty"`

	// metadata
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}
