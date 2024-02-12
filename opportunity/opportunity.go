package opportunity

import (
	"github.com/FrosTiK-SD/models/misc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Opportunity struct {
	ID                     primitive.ObjectID      `bson:"_id" json:"_id,omitempty"`
	Company                primitive.ObjectID      `bson:"company" json:"company,omitempty"`
	Profiles               []primitive.ObjectID    `bson:"profiles" json:"profiles,omitempty"`
	DetailsRequestedSchema *map[string]interface{} `bson:"detailsRequestedSchema" json:"detailsRequestedSchema"`
	Attachments            []misc.Attachment       `bson:"attachments" json:"attachments,omitempty"`
	Extras                 *string                 `bson:"extras" json:"extras,omitempty"`
	MaxAllowedProfiles     int                     `bson:"maxAllowedProfiles" json:"maxAllowedProfiles"`

	// metadata
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}
