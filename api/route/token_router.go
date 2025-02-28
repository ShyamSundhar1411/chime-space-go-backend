package route

import (
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/api/controller"
	"github.com/ShyamSundhar1411/chime-space-go-backend/bootstrap"
	"github.com/ShyamSundhar1411/chime-space-go-backend/usecase"
	"github.com/labstack/echo/v4"
)

func NewTokenRouter(env *bootstrap.Env,timeout time.Duration,routerGroup echo.Group){
	tokenController := &controller.TokenController{
		TokenUseCase: usecase.NewTokenUseCase(timeout),
		Env: env,
	}
	routerGroup.POST("/token", tokenController.Refresh)
}