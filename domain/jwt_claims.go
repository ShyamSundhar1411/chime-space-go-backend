package domain

import "github.com/golang-jwt/jwt/v5"

type CustomJWTClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.Claims
}

type CustomJWTRefreshClaims struct {
	ID string `json:"id"`
	jwt.Claims
}