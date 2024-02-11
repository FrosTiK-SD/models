package misc

import "go.mongodb.org/mongo-driver/bson/primitive"

type ActivityType string

const (
	SEEN    ActivityType = "SEEN"
	CREATED ActivityType = "CREATE"
	EDITED  ActivityType = "EDIT"
	DELETED ActivityType = "DELETE"
)

type Activity struct {
	Type      ActivityType       `json:"type" bson:"type"`
	Timestamp primitive.DateTime `json:"timestamp" bson:"timestamp"`
	User      primitive.ObjectID `json:"user" bson:"user"`
}
