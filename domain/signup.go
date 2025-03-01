package domain

import (
	"context"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

type SignUpRequest struct {
	UserName string `form:"userName" binding:"required"`
	PenName  string `form:"penName" binding:"reqired"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type SignUpResponse struct {
	Message      string `json:"message"`
	StatusCode   int    `json:"statusCode"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User         models.User   `json:"user"`
}

type SignUpUsecase interface {
	Create(c context.Context, user *models.User) error
	GetUserByUsername(c context.Context, username string) (models.User, error)
	CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *models.User, secret string, exprity int) (refreshToken string, err error)
}
