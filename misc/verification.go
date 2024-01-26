package misc

import "go.mongodb.org/mongo-driver/bson/primitive"

type Verification struct {
	IsVerified bool               `bson:"isVerified" json:"isVerified,omitempty"`
	VerifiedBy primitive.ObjectID `bson:"verifiedBy" json:"verifiedBy,omitempty"`
	VerifiedAt primitive.DateTime `bson:"verifiedAt" json:"verifiedAt,omitempty"`
}
