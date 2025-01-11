package entities

import (
	"context"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}
type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type LoginUsecase interface {
	GetUserByUserName(c context.Context,userName string)(models.User,error)
	CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error)
}