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

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, publicRouterGroup *echo.Group, privateRouterGroup *echo.Group){
	userRepository := repository.NewUserRepository(db, domain.CollectionUser)
	userController := &controller.UserController{
		UserUseCase: usecase.NewUserUseCase(userRepository, timeout),
		Env: env,
	}
	privateRouterGroup.GET("/user/me/", userController.GetMyProfile)

}