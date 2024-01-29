package exam

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coordinates struct {
	Row    int32 `json:"row" bson:"row"`
	Column int32 `json:"column" bson:"column"`
}

type MarkData struct {
	User primitive.ObjectID `json:"user" bson:"user"`
	Time string             `json:"time" bson:"time"`
}

type Attendance struct {
	Id            primitive.ObjectID `json:"_id" bson:"_id"`
	ExaminationId primitive.ObjectID `json:"examinationId" bson:"examinationId"`
	User          primitive.ObjectID `json:"user" bson:"user"`
	Venue         primitive.ObjectID `json:"venue" bson:"venue"`
	RawPosition   Coordinates        `json:"rawPosition" bson:"rawPosition"`
	IsPresent     bool               `json:"isPresent" bson:"isPresent"`
	MarkedData    []MarkData         `json:"markedData" bson:"markedData"`
}
