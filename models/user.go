package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID       bson.ObjectID `bson:"_id" json:"id"`
	UserName string        `bson:"username" json:"userName"`
	PenName  string        `bson:"penname" json:"penName"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password,omitempty" json:"-"`
}