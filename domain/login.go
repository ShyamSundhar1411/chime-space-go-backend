package domain

import (
	"context"
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
	User         User   `json:"user"`
}
type LoginUsecase interface {
	GetUserByUserName(c context.Context, userName string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
