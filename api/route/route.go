package route

import (
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/bootstrap"
	"github.com/ShyamSundhar1411/chime-space-go-backend/mongo"
	"github.com/labstack/echo/v4"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, echoEngine *echo.Echo) {
	publicRouter := echoEngine.Group("")
	publicRouter.POST("/", func(c echo.Context) error {
		return c.String(200, "Hello World")
	})

}
