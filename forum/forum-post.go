package forum

import (
	company "github.com/FrosTiK-SD/models/company"
	"github.com/FrosTiK-SD/models/constant"
	"github.com/FrosTiK-SD/models/student"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyProfile struct {
	Company company.Company `bson:"company,omitempty" json:"company,omitempty"`
	IAF     company.IAF     `bson:"iaf,omitempty" json:"iaf,omitempty"`
	JAF     company.JAF     `bson:"jaf,omitempty" json:"jaf,omitempty"`
}

type ForumGroup struct {
	Branch constant.Branch `bson:"branch" json:"branch,omitempty"`
	Course constant.Course `bson:"course" json:"course,omitempty"`
	Batch  student.Batch   `bson:"batch" json:"batch,omitempty"`
}

type NotifyForum struct {
	StudentList []primitive.ObjectID `bson:"studentList,omitempty" json:"studentList"`
	GroupsForum []string             `bson:"groupsForum,omitempty" json:"groupsForum"`
}

type ForumPost struct {
	ID               primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Title            string               `bson:"title,omitempty" json:"title,omitempty"`
	Type             ForumPostType        `bson:"type,omitempty" json:"type,omitempty"`
	Tags             []string             `bson:"tags,omitempty" json:"tags,omitempty"`
	RelatedCompanies []primitive.ObjectID `bson:"relatedCompanies,omitempty" json:"relatedCompanies,omitempty"`
	Notify           NotifyForum          `bson:"notify,omitempty" json:"notify,omitempty"`
	Content          []ForumPostContent   `bson:"content,omitempty" json:"content,omitempty"`
	Pinned           int                  `bson:"pinned,omitempty" json:"pinned,omitempty"`

	// metadata
	CreatedBy primitive.ObjectID   `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	EditedBy  []primitive.ObjectID `bson:"editedBy,omitempty" json:"editedBy,omitempty"`
	CreatedAt primitive.DateTime   `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime   `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	RawText   string               `bson:"rawText,omitempty" json:"rawText,omitempty"`
}

type ForumPostContent struct {
	TopGenericContent    string               `bson:"topGenericContent,omitempty" json:"topGenericContent,omitempty"`
	BottomGenericContent string               `bson:"bottomGenericContent,omitempty" json:"bottomGenericContent,omitempty"`
	StudentList          []primitive.ObjectID `bson:"studentList,omitempty" json:"studentList,omitempty"`
}
