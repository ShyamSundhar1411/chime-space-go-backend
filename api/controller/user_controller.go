package controller

import (
	"net/http"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"github.com/labstack/echo/v4"
)

func (userController *UserController) GetMyProfile(c echo.Context) error {
	ctx := utils.ExtractContext(c)
	user, err := userController.UserUseCase.GetMyProfile(ctx)
	if err != nil{
		return c.JSON(http.StatusNotFound, domain.ProfileResponse{
			Message: "Profile Not Found",
			StatusCode: 404,
			Profile: nil,
		})
	}
	return c.JSON(http.StatusOK, domain.ProfileResponse{
		Message: "Profile Found",
		StatusCode: 200,
		Profile: user,
	})

}