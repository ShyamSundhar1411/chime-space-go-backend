package domain

import "context"

type TokenRefeshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenRefreshResponse struct {
	AccessToken string `json:"access_token"`
}

type TokenUsecase interface {
	RefreshToken(c context.Context, refreshToken string) (string, error)
}