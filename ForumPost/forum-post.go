package forumpost

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	company "github.com/FrosTiK-SD/models-go/company"
)

type CompanyProfile {
	Company company.Company
}

type ForumPost struct {
	ID               primitive.ObjectID
	Title            string
	Type             string
	Tags             []string
	RelatedCompanies CompanyProfile
}
