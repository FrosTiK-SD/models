package opportunity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Slot struct {
	Day                 primitive.DateTime `json:"day,omitempty"`
	StartDate           primitive.DateTime `json:"start_date,omitempty"`
	EndDate             primitive.DateTime `json:"end_date,omitempty"`
	Name                string             `json:"name,omitempty"`
	Number              int                `json:"number,omitempty"`         // slot number
	MaxInterviews       int                `json:"max_interviews,omitempty"` // MaximumNumber of Interviews
	PreferenceStartTime primitive.DateTime `json:"preference_start_time,omitempty"`
	PreferenceEndTime   primitive.DateTime `json:"preference_end_time,omitempty"`
}
