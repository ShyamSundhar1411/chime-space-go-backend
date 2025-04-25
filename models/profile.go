package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Profile struct {
	Bio            string        `bson:"bio" json:"bio"`
	ProfilePicture string        `bson:"profilePicture" json:"profilePicture"`
	CoverPicture   string        `bson:"coverPicture" json:"coverPicture"`
	Followers      []string      `bson:"followers" json:"followers"`
	Following      []string      `bson:"following" json:"following"`
	User           bson.ObjectID `bson:"user" json:"user"`
}
