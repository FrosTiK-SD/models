package willingness

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShortList struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Profile  primitive.ObjectID   `bson:"profile,omitempty"`
	Type     string               `bson:"type,omitempty"`
	Students []primitive.ObjectID `bson:"student_list,omitempty"`

	//metadata
	CreatedBy primitive.ObjectID   `bson:"created_by,omitempty" json:"created_by,omitempty"`
	EditedBy  []primitive.ObjectID `bson:"edited_by,omitempty" json:"edited_by,omitempty"`
	CreatedAt primitive.DateTime   `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt primitive.DateTime   `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
