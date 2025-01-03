package utils

import (
	"context"

	"github.com/labstack/echo/v4"
)

func ExtractContext(c echo.Context) context.Context {
	return c.Request().Context()
}
