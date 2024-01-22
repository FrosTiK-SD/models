package forum

import (
	company "github.com/FrosTiK-SD/frostik-models/company"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyProfile struct {
	Company company.Company `bson:"company,omitempty" json:"company,omitempty"`
	IAF     company.IAF     `bson:"iaf,omitempty" json:"iaf,omitempty"`
	JAF     company.JAF     `bson:"jaf,omitempty" json:"jaf,omitempty"`
}

type NotificationGroups struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NotifyTo NotifyToType       `bson:"notifyTo,omitempty" json:"notifyTo,omitempty"`
}

type NotifyForum struct {
	StudentList []primitive.ObjectID `bson:"studentList,omitempty" json:"studentList"`
	GroupsForum []primitive.ObjectID `bson:"groupsForum,omitempty" json:"grouptsForum"`
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
}

type ForumPostContent struct {
	TopGenericContent    map[string]interface{} `bson:"topGenericContent,omitempty" json:"topGenericContent,omitempty"`
	BottomGenericContent map[string]interface{} `bson:"bottomGenericContent,omitempty" json:"bottomGenericContent,omitempty"`
	StudentList          []primitive.ObjectID   `bson:"studentList,omitempty" json:"studentList,omitempty"`
}
