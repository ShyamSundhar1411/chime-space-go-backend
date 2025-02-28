package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary      Generate a new access token using a refresh token
// @Description  Allows users to refresh their access token using a valid refresh token.
// @Tags         Token
// @Accept       json
// @Produce      json
// @Param        refreshRequest  body  domain.TokenRefreshRequest  true  "Refresh Token Request Payload"
// @Success      200  {object}  domain.TokenRefreshResponse  "Returns a new access token"
// @Router       /token/refresh/ [post]
func (tokenController *TokenController) Refresh(c echo.Context) error {
	return c.JSON(http.StatusOK, "Testing endpoint")
}
