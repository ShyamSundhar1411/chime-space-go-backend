package utils

import (
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &domain.CustomJWTClaims{
		Name: user.UserName,
		ID:   user.ID.Hex(),
		Claims: jwt.MapClaims{
			"exp": exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	refreshClaims := &domain.CustomJWTRefreshClaims{
		ID: user.ID.Hex(),
		Claims: jwt.MapClaims{
			"exp": exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func VerifyRefreshToken(tokenString string, secret string) (*domain.CustomJWTRefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.CustomJWTRefreshClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*domain.CustomJWTRefreshClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
