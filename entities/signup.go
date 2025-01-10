package entities

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
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignUpUsecase interface {
	Create(c context.Context, user *models.User)error
	GetUserByEmail(c context.Context, email string)(models.User, error)
	CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *models.User,secret string,exprity int)(refreshToken string,err error)
}