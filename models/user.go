package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID primitive.ObjectID `bson:"_id"`
	Username string `bson:"username"`
	Penname string `bson:"penname"`
	Email string `bson:"email"`
	Password string `bson:"password"`
}

type UserRepository interface{
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetById(c context.Context, id string) (User, error)
	GetByUsername(c context.Context, username string) (User, error)
}