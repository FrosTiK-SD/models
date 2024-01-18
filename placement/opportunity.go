package opportunity

import (
	"github.com/FrosTiK-SD/models-go/company"
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
	ID                        primitive.ObjectID          `json:"id,omitempty"`
	Company                   primitive.ObjectID          `json:"company,omitempty"`
	Profile                   string                      `json:"profile,omitempty"`
	Criteria                  company.EligibilityCriteria `json:"criteria,omitempty"`
	DetailsRequestedGenerator map[string]interface{}      `json:"details_requested_generator,omitempty"` // formbuilder

	OpportunityStartTime primitive.DateTime `json:"opportunity_start_time,omitempty"`
	OpportunityEndTime   primitive.DateTime `json:"opportunity_end_time,omitempty"`
	Mode                 Listing            `json:"mode,omitempty"`
	OpportunityType      OpportunityType    `json:"opportunity_type,omitempty"`
}

type OpportunityPopulated struct {
	ID                        primitive.ObjectID          `json:"id,omitempty"`
	Company                   company.Company             `json:"company,omitempty"`
	Profile                   string                      `json:"profile,omitempty"`
	Criteria                  company.EligibilityCriteria `json:"criteria,omitempty"`
	DetailsRequestedGenerator map[string]interface{}      `json:"details_requested_generator,omitempty"` // formbuilder

	OpportunityStartTime primitive.DateTime `json:"opportunity_start_time,omitempty"`
	OpportunityEndTime   primitive.DateTime `json:"opportunity_end_time,omitempty"`
	Mode                 Listing            `json:"mode,omitempty"`
	OpportunityType      OpportunityType    `json:"opportunity_type,omitempty"`
}
