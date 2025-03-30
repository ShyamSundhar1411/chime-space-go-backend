package domain

import (
	"context"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

type TokenRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenRefreshResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenUsecase interface {
	GenerateAccessToken(c context.Context, user *models.User, secret string, expiry int) (string, error)
	GenerateRefreshToken(c context.Context, user *models.User, secret string, expiry int)(string, error)
	ValidateRefreshToken(c context.Context, refreshToken string, secret string) (*models.User, error)
}