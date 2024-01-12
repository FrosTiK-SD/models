package placement

import "go.mongodb.org/mongo-driver/bson/primitive"

type Slot struct {
	Day                 primitive.DateTime
	StartTime           primitive.DateTime
	EndTime             primitive.DateTime
	Name                string
	Number              int // slot number
	MaxInterviews       int // MaximumNumber of Interviews
	PreferenceStartTime primitive.DateTime
	PreferenceEndTime   primitive.DateTime
}
