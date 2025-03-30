package domain

import (
	"context"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

type TokenRefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type TokenRefreshResponse struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	StatusCode int `json:"statusCode"`
	Message		string `json:"message"`
}

type TokenUsecase interface {
	GenerateAccessToken(c context.Context, user *models.User, secret string, expiry int) (string, error)
	GenerateRefreshToken(c context.Context, user *models.User, secret string, expiry int)(string, error)
	ValidateRefreshToken(c context.Context, refreshToken string, secret string) (*models.User, error)
}