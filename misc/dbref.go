package misc

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBRef struct {
	Ref string             `bson:"$ref" json:"$ref"`
	ID  primitive.ObjectID `bson:"$id" json:"$id"`
	DB  string             `bson:"$db" json:"$db"`
}
