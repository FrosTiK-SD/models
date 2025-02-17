package exam

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	constant "github.com/FrosTiK-SD/models/constant"

)

type Examination struct {
	Id          primitive.ObjectID   `json:"_id" bson:"_id"`
	Title       string               `json:"title" bson:"title"`
	Description string               `json:"description" bson:"description"`
	LinkedIaf   []primitive.ObjectID `json:"linkedIaf" bson:"linkedIaf"`
	LinkedJaf   []primitive.ObjectID `json:"linkedJaf" bson:"linkedJaf"`
	Venues      []primitive.ObjectID `json:"venues" bson:"venues"`
	Attendance  []primitive.ObjectID `json:"attendance" bson:"attendance"`
	StartDate   string               `json:"startDate" bson:"startDate"`
	EndDate     string               `json:"endDate" bson:"endDate"`
	StartBuffer time.Duration        `json:"startBuffer" bson:"startBuffer"`
	EndBuffer   time.Duration        `json:"endBuffer" bson:"endBuffer"`
	CreatedBy   primitive.ObjectID   `json:"createdBy" bson:"createdBy"`
	WillingList []primitive.ObjectID `json:"willingList" bson:"willingList"`
	IsActive    bool                 `json:"isActive" bson:"isActive"`
	Session    constant.AcademicSession `json:"session" bson:"session"`
}
