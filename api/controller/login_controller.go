package controller

import (
	"net/http"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// @Summary		User Login
// @Description	Logs a user in by validating credentials and returning access and refresh tokens.
// @Tags			Authentication
// @Accept			json
// @Produce		json
// @Param			loginRequest	body		domain.LoginRequest		true	"Login Request Payload"
// @Success		200				{object}	domain.LoginResponse	"Login successful, returns access and refresh tokens"
// @Failure		400				{object}	domain.ErrorResponse	"Invalid request"
// @Failure		401				{object}	domain.ErrorResponse	"Unauthorized - Invalid credentials or user not found"
// @Failure		500				{object}	domain.ErrorResponse	"Internal Server Error"
// @Router			/auth/login/ [post]
func (loginController *LoginController) Login(c echo.Context) error {
	var request domain.LoginRequest
	err := c.Bind(&request)
	ctx := utils.ExtractContext(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid request", StatusCode: http.StatusBadRequest})
	}
	user, err := loginController.LoginUsecase.GetUserByUserName(ctx, request.UserName)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found with given credentials", StatusCode: http.StatusUnauthorized})
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials", StatusCode: http.StatusUnauthorized})
	}
	accessToken, err := loginController.LoginUsecase.CreateAccessToken(&user, loginController.Env.AccessTokenSecretKey, loginController.Env.AccessTokenExpiryHour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error(), StatusCode: http.StatusInternalServerError})
	}
	refreshToken, err := loginController.LoginUsecase.CreateRefreshToken(&user, loginController.Env.RefreshTokenSecretKey, loginController.Env.RefreshTokenExpiryHour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error(), StatusCode: http.StatusInternalServerError})

	}

	loginResponse := domain.LoginResponse{
		StatusCode:   http.StatusOK,
		Message:      "Login successful",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}
	return c.JSON(http.StatusOK, loginResponse)

}
