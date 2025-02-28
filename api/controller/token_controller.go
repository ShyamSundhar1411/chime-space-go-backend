package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (tokenController *TokenController) Refresh(c echo.Context) error{
	return c.JSON(http.StatusOK, "Testing endpoint")
}