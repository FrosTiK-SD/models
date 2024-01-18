package resume

import "go.mongodb.org/mongo-driver/bson/primitive"

type Resume struct {
	ID      primitive.ObjectID
	Student primitive.ObjectID
	Title   string
}
