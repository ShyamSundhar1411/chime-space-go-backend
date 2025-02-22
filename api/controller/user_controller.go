package controller

import (
	"net/http"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"github.com/labstack/echo/v4"
)
// @Summary		User Me Endpoint
// @Description	Retrives the profile details of the user based on access token
// @Tags			User
// @Accept			json
// @Produce		json
// @Success		201				{object}	domain.ProfileResponse	"Returns Profile of the user"
// @Failure		404				{object}	domain.BaseResponse	"Profile Not Found"
// @Router			/user/me/	 [get]
//	@Security		BearerAuth
func (userController *UserController) GetMyProfile(c echo.Context) error {
	ctx := utils.ExtractContext(c)
	user, err := userController.UserUseCase.GetMyProfile(ctx)
	if err != nil{
		return c.JSON(http.StatusNotFound, domain.BaseResponse{
			Message: "Profile Not Found",
			StatusCode: http.StatusNotFound,

		})
	}
	return c.JSON(http.StatusOK, domain.ProfileResponse{
		Message: "Profile Found",
		StatusCode: http.StatusOK,
		Profile: user,
	})
}