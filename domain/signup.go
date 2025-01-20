package domain

import (
	"context"
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
	User         User   `json:"user"`
}

type SignUpUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByUsername(c context.Context, username string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, exprity int) (refreshToken string, err error)
}
