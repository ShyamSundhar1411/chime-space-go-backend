package domain

import (
	"context"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

type LoginRequest struct {
	UserName string `form:"username" binding:"required,username"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	Message      string `json:"message"`
	StatusCode   int    `json:"statusCode"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User         models.User   `json:"user"`
}
type LoginUsecase interface {
	GetUserByUserName(c context.Context, userName string) (models.User, error)
	CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error)
}
