package opportunity

import (
	"github.com/FrosTiK-SD/models/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Criteria struct {
	Id             primitive.ObjectID `bson:"_id" json:"_id"`
	Name           string             `bson:"name" json:"name"`
	X              float64            `bson:"x" json:"x"`
	XII            float64            `bson:"xii" json:"xii"`
	Cgpa           float64            `bson:"cgpa" json:"cgpa"`
	Branches       []constant.Branch  `bson:"branches" json:"branches"`
	ActiveBacklogs *int               `bson:"activeBacklogs" json:"activeBacklogs"`
	TotalBacklogs  *int               `bson:"totalBacklogs" json:"totalBacklogs"`
	Gender         []constant.Gender  `bson:"gender" json:"gender"`
	Disbility      []bool             `bson:"disability" json:"disability"`
	Courses        []constant.Course  `bson:"courses" json:"courses" validate:"required"`
	CreatedAt      primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt      primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}
