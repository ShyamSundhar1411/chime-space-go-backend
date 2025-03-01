package domain

import "context"

type TokenRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenRefreshResponse struct {
	AccessToken string `json:"access_token"`
}

type TokenUsecase interface {
	RefreshToken(c context.Context, refreshToken string) (string, error)
	ValidateRefreshToken(c context.Context, refreshToken string, secret string) (bool, error)
}