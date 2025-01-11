package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID       bson.ObjectID `bson:"_id"`
	UserName string        `bson:"username"`
	PenName  string        `bson:"penname"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetById(c context.Context, id string) (User, error)
	GetByUsername(c context.Context, username string) (User, error)
}
