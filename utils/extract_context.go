package utils

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func ExtractContext(c echo.Context) context.Context {
	ctx := c.Request().Context()
	user := c.Get("user")
	if user == nil {
		return ctx
	}
	token, ok := user.(*jwt.Token)
	if !ok {
		return ctx
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {

		return ctx
	}
	userId, ok := claims["id"].(string)
	if !ok || userId == "" {
		return ctx
	}
	ctx = context.WithValue(ctx, "userId", userId)
	return ctx
}
