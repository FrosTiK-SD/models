package placement

import (
	"github.com/FrosTiK-SD/models-go/company"
	"github.com/FrosTiK-SD/models-go/student"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Listing string
type OpportunityType string
type Gender string

const (
	LISTED             Listing = "listed"
	UNDER_BRANCH_ISSUE Listing = "under_branch_issue"
)

const (
	PLACEMENT OpportunityType = "placement"
	INTERN    OpportunityType = "intern"
)

const (
	MALE         Gender = "male"
	FEMALE       Gender = "female"
	OTHER_GENDER Gender = "other"
)

type Opportunity struct {
	Company         company.Company
	Profile         string
	Criteria        company.EligibilityCriteria
	DetailRequested map[string]interface{} // formbuilder

	OpportunityStartTime primitive.DateTime
	OpportunityEndTime   primitive.DateTime
	Mode                 Listing
	OpportunityType      OpportunityType
}

type Application struct {
	Opportunity Opportunity
	Student     student.Student
	// Resume           Resume
	DetailsRequested map[string]interface{}
}

type Criteria struct {
	Xth           float32
	Xiith         float32
	CGPA          float32
	Branch        []string
	Course        []string
	ActiveBacklog int
	TotalBacklog  int
	Gender        Gender
	Disability    bool
}
