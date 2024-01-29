package exam

import "go.mongodb.org/mongo-driver/bson/primitive"

type Venue struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Arrangement string             `json:"arrangement" bson:"arrangement"`
	Map         [][]int32          `json:"map" bson:"map"`
	Capacity    int32              `json:"capacity" bson:"capacity"`
}
