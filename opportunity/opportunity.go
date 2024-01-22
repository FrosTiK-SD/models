package opportunity

import (
	"github.com/FrosTiK-SD/iitbhu-tpc-models-golang/misc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Opportunity struct {
	ID                     primitive.ObjectID     `bson:"id" json:"id,omitempty"`
	Company                primitive.ObjectID     `bson:"company" json:"company,omitempty"`
	Profiles               []primitive.ObjectID   `bson:"profile" json:"profile,omitempty"`
	DetailsRequestedSchema map[string]interface{} `bson:"detailsRequestedSchema" json:"detailsRequestedSchema"`
	Attachments            []misc.Attachment      `bson:"attachments" json:"attachments,omitempty"`

	// metadata
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}