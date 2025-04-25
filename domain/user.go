package domain

import (
	"context"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

const (
	CollectionUser = "users"
)

type UserProfile struct {
	User    *models.User
	Profile *models.Profile
}
type UserUsecase interface {
	GetMyProfile(c context.Context) (*UserProfile, error)
}
type UserController interface {
	GetMyProfile(c context.Context) (UserProfile, error)
}

type UserProfileResponse struct {
	Message    string       `json:"message"`
	StatusCode int          `json:"statusCode"`
	Profile    *UserProfile `json:"profile"`
}
type UserRepository interface {
	Create(c context.Context, user *models.User) error
	Fetch(c context.Context) ([]models.User, error)
	GetById(c context.Context, id string) (*models.User, error)
	GetByUsername(c context.Context, username string) (models.User, error)
	GetMyProfile(c context.Context) (*models.User, error)
}
