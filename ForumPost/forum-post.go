package forumpost

import (
	company "github.com/FrosTiK-SD/models-go/company"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyProfile struct {
	Company company.Company
}

type ForumPost struct {
	ID               primitive.ObjectID
	Title            string
	Type             string
	Tags             []string
	RelatedCompanies CompanyProfile
}
