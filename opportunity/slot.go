package opportunity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Slot struct {
	ID                  primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Day                 primitive.DateTime `bson:"day" json:"day,omitempty"`
	StartDate           primitive.DateTime `bson:"start_date" json:"start_date,omitempty"`
	EndDate             primitive.DateTime `bson:"end_date" json:"end_date,omitempty"`
	SlotName            string             `bson:"name" json:"name,omitempty"`
	SlotNumber          int                `bson:"number" json:"number,omitempty"`
	MaxInterviews       int                `bson:"max_interviews" json:"max_interviews,omitempty"`
	PreferenceStartTime primitive.DateTime `bson:"preference_start_time" json:"preference_start_time,omitempty"`
	PreferenceEndTime   primitive.DateTime `bson:"preference_end_time" json:"preference_end_time,omitempty"`

	// metadata
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updatedAt" json:"updatedAt,omitempty"`
}
