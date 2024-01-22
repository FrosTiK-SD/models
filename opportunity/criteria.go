package opportunity

import (
	"github.com/FrosTiK-SD/iitbhu-tpc-models-golang/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Criteria struct {
	ID            primitive.ObjectID `bson:"id" json:"id,omitempty"`
	Xth           float32            `bson:"xth" json:"xth,omitempty"`
	XIIth         float32            `bson:"xiith" json:"xiith,omitempty"`
	CGPA          float32            `bson:"cgpa" json:"cgpa,omitempty"`
	Branches      []constant.Branch  `bson:"branches" json:"branches,omitempty"`
	Courses       []constant.Course  `bson:"courses" json:"courses,omitempty"`
	ActiveBacklog int                `bson:"active_backlog" json:"active_backlog,omitempty"`
	TotalBacklog  int                `bson:"total_backlog" json:"total_backlog,omitempty"`
	Gender        constant.Gender    `bson:"gender" json:"gender,omitempty"`
	Disability    bool               `bson:"disability" json:"disability,omitempty"`

	// metadata
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}
