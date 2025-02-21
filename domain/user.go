package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID       bson.ObjectID `bson:"_id" json:"id"`
	UserName string        `bson:"username" json:"userName"`
	PenName  string        `bson:"penname" json:"penName"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password,omitempty" json:"-"`
}
type UserUsecase interface {
	GetMyProfile(c context.Context)(*User, error)
}
type UserController interface {
	GetMyProfile(c context.Context)(*User, error)
}
type ProfileResponse struct {
	Message      string `json:"message"`
	StatusCode   int    `json:"statusCode"`
	Profile      *User   `json:"profile"`
}
type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetById(c context.Context, id string) (User, error)
	GetByUsername(c context.Context, username string) (User, error)
	GetMyProfile(c context.Context)(*User,error)
}
