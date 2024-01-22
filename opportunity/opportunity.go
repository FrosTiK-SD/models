package opportunity

import (
	"github.com/FrosTiK-SD/iitbhu-tpc-models-golang/misc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OpportunityType string

const (
	MULTIPLE_PROFILE_SINGLE_WILLINGNESS   OpportunityType = "MULTIPLE_PROFILE_SINGLE_WILLINGNESS"
	MULTIPLE_PROFILE_MULTIPLE_WILLINGNESS OpportunityType = "MULTIPLE_PROFILE_MULTIPLE_WILLINGNESS"
)

type Opportunity struct {
	ID                     primitive.ObjectID     `bson:"id" json:"id,omitempty"`
	Company                primitive.ObjectID     `bson:"company" json:"company,omitempty"`
	Profiles               []primitive.ObjectID   `bson:"profile" json:"profile,omitempty"`
	DetailsRequestedSchema map[string]interface{} `bson:"detailsRequestedSchema" json:"detailsRequestedSchema"`
	Attachments            []misc.Attachment      `bson:"attachments" json:"attachments,omitempty"`
	Type                   OpportunityType        `bson:"type" json:"type,omitempty"`

	// metadata
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}
