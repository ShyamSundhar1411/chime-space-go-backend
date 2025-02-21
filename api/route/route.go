package route

import (
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/bootstrap"
	"github.com/ShyamSundhar1411/chime-space-go-backend/mongo"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, echoEngine *echo.Echo) {
	publicRouter := echoEngine.Group("")

	NewLoginRouter(env, timeout, db, publicRouter)
	NewSignUpRouter(env, timeout, db, publicRouter)
	privateRouter := echoEngine.Group("")
	config := echojwt.Config{
		SigningKey: []byte(env.AccessTokenSecretKey),
	}
	privateRouter.Use(echojwt.WithConfig(config))
	NewChimeRouter(env, timeout, db, publicRouter, privateRouter)
	NewUserRouter(env, timeout, db, publicRouter, privateRouter)
}
