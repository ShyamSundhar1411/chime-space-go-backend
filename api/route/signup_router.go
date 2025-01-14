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

func NewSignUpRouter(env *bootstrap.Env, timout time.Duration, db mongo.Database, routerGroup *echo.Group) {
	userRepository := repository.NewUserRepository(db, domain.CollectionUser)
	signUpController := &controller.SignUpController{
		SignUpUsecase: usecase.NewSignUpUsecase(userRepository, timout),
		Env:           env,
	}
	routerGroup.POST("/auth/signup/", signUpController.SignUp)

}
