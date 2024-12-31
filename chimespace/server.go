package main

import (
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/api/route"
	"github.com/ShyamSundhar1411/chime-space-go-backend/bootstrap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.DBName)

	defer app.CloseDBConnection()
	timeout := time.Duration(env.ContextTimeout) * time.Second
	e := echo.New()
	e.Logger.Info(db)
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Request Timeout",

		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			e.Logger.Fatal(c.Path())
		},
		Timeout: timeout,
	}))
	route.Setup(env, timeout,db,e)
	e.Logger.Fatal(e.Start(":8080"))

}
