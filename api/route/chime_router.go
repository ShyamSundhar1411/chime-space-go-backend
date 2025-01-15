package route

import (
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/api/controller"
	"github.com/ShyamSundhar1411/chime-space-go-backend/bootstrap"
	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/mongo"
	"github.com/ShyamSundhar1411/chime-space-go-backend/repository"
	"github.com/ShyamSundhar1411/chime-space-go-backend/usecase"
	"github.com/labstack/echo/v4"
)

func NewChimeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, publicRouterGroup *echo.Group,privateRouterGroup *echo.Group) {
	chimeRepository := repository.NewChimeRepository(db, domain.CollectionChime)
	chimeController := &controller.ChimeController{
		ChimeUsecase: usecase.NewChimeUseCase(
			chimeRepository,
			timeout,
		),
	}
	publicRouterGroup.GET("/chimes/", chimeController.FetchAllChimes)
	privateRouterGroup.GET("/chimes/user/",chimeController.FetchChimeFromUser)
}
