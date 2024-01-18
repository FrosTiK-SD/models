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
	NotifyTo NotifyToType       `bson:"notify_to,omitempty" json:"notify_to,omitempty"`
}

type NotifyForum struct {
	StudentList []primitive.ObjectID `bson:"student_list,omitempty"`
	GroupsForum []primitive.ObjectID `bson:"groups_forum,omitempty"`
}

type ForumPost struct {
	ID               primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Title            string               `bson:"title,omitempty" json:"title,omitempty"`
	Type             ForumPostType        `bson:"type,omitempty" json:"type,omitempty"`
	Tags             []string             `bson:"tags,omitempty" json:"tags,omitempty"`
	RelatedCompanies []primitive.ObjectID `bson:"related_companies,omitempty" json:"related_companies,omitempty"`
	Notify           NotifyForum          `bson:"notify,omitempty" json:"notify,omitempty"`
	Content          []ForumPostContent   `bson:"content,omitempty" json:"content,omitempty"`
	Pinned           int                  `bson:"pinned,omitempty" json:"pinned,omitempty"`

	// metadata
	CreatedBy primitive.ObjectID   `bson:"created_by,omitempty" json:"created_by,omitempty"`
	EditedBy  []primitive.ObjectID `bson:"edited_by,omitempty" json:"edited_by,omitempty"`
	CreatedAt primitive.DateTime   `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt primitive.DateTime   `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type ForumPostContent struct {
	TopGenericContent    map[string]interface{} `bson:"top_generic_content,omitempty" json:"top_generic_content,omitempty"`
	BottomGenericContent map[string]interface{} `bson:"bottom_generic_content,omitempty" json:"bottom_generic_content,omitempty"`
	StudentList          []primitive.ObjectID   `bson:"student_list" json:"student_list"`
}
