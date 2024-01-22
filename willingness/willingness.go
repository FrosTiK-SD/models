package willingness

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShortList struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Profile  primitive.ObjectID   `bson:"profile,omitempty" json:"profile,omitempty"`
	Type     string               `bson:"type,omitempty" json:"type,omitempty"`
	Students []primitive.ObjectID `bson:"studentList,omitempty" json:"studentList,omitempty"`

	//metadata
	CreatedBy primitive.ObjectID   `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	EditedBy  []primitive.ObjectID `bson:"editedBy,omitempty" json:"editedBy,omitempty"`
	CreatedAt primitive.DateTime   `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime   `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
