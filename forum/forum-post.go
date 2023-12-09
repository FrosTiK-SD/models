package forum

import (
	company "github.com/FrosTiK-SD/models-go/company"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyProfile struct {
	Company company.Company
	IAF     company.IAF
	JAF     company.JAF
}

type NotificationGroups struct {
	ID   primitive.ObjectID
	Type string
}

type ForumPost struct {
	ID               primitive.ObjectID
	Title            string
	Type             string
	Tags             []string
	RelatedCompanies []CompanyProfile
	RelatedStudents  []primitive.ObjectID
	Content          ForumPostContent
	Pinned           int

	// metadata
	CreatedBy primitive.ObjectID
	EditedBy  []primitive.ObjectID
	CreatedAt primitive.DateTime
	UpdatedAt primitive.DateTime
}

type ForumPostContent struct {
	Type ForumPostContentType
	Data interface{}
}
