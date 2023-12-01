package company

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	PRIVATE = "private"
	PSU     = "Public Sector Unit"
)

type Company struct {
	ID        primitive.ObjectID
	Name      string
	Logo      []string
	Address   []string
	Category  string
	Sector    string
	Website   string
	CreatedAt primitive.DateTime
	UpdatedAt primitive.DateTime
}
