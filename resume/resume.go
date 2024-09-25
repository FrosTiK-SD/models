package resume

import "go.mongodb.org/mongo-driver/bson/primitive"

type Resume struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Student    primitive.ObjectID `json:"student" bson:"student" binding:"required"`
	Name       string             `json:"name" bson:"name" binding:"required"`
	PDF        ResumePdf          `json:"pdf" bson:"pdf" binding:"required"`
	IsVerified bool               `json:"isVerified" bson:"isVerified" binding:"required"`
}

type ResumePdf struct {
	Url   string `json:"url,omitempty" bson:"url,omitempty"`
	Latex string `json:"latex,omitempty" bson:"latex,omitempty"`
}
