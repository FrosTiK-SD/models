package misc

import "go.mongodb.org/mongo-driver/bson/primitive"

type Verification struct {
	IsVerified bool               `bson:"isVerified" json:"isVerified"`
	VerifiedBy primitive.ObjectID `bson:"verifiedBy" json:"verifiedBy"`
	VerifiedAt primitive.DateTime `bson:"verifiedAt" json:"verifiedAt"`
}
