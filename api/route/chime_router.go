package route

import (
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/api/controller"
	"github.com/ShyamSundhar1411/chime-space-go-backend/bootstrap"
	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
	"github.com/ShyamSundhar1411/chime-space-go-backend/mongo"
	"github.com/ShyamSundhar1411/chime-space-go-backend/repository"
	"github.com/ShyamSundhar1411/chime-space-go-backend/usecase"
	"github.com/labstack/echo/v4"
)

func NewChimeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, routerGroup *echo.Group) {
	chimeRepository := repository.NewChimeRepository(db, models.CollectionChime)
	chimeController := &controller.ChimeController{
		ChimeUsecase: usecase.NewChimeUseCase(
			chimeRepository,
			timeout,
		),
	}
	routerGroup.GET("/chimes", chimeController.FetchAllChimes)
	routerGroup.POST("/chime", func(c echo.Context) error {
		return c.String(201, "Posted, Chimes!")
	})
}
