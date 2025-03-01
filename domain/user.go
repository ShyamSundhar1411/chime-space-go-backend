package domain

import (
	"context"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

const (
	CollectionUser = "users"
)


type UserUsecase interface {
	GetMyProfile(c context.Context)(*models.User, error)
}
type UserController interface {
	GetMyProfile(c context.Context)(*models.User, error)
}
type ProfileResponse struct {
	Message      string `json:"message"`
	StatusCode   int    `json:"statusCode"`
	Profile      *models.User   `json:"profile"`
}
type UserRepository interface {
	Create(c context.Context, user *models.User) error
	Fetch(c context.Context) ([]models.User, error)
	GetById(c context.Context, id string) (models.User, error)
	GetByUsername(c context.Context, username string) (models.User, error)
	GetMyProfile(c context.Context)(*models.User,error)
}
