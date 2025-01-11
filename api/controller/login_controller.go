package controller

import (
	"net/http"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (loginController *LoginController) Login(c echo.Context){
	var request domain.LoginRequest
	err := c.Bind(&request)
	ctx := utils.ExtractContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest,domain.ErrorResponse{Message: "Invalid request",StatusCode: http.StatusBadRequest})
	}
	user, err := loginController.LoginUsecase.GetUserByUserName(ctx, request.UserName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found with given credentials", StatusCode: http.StatusUnauthorized})
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials", StatusCode: http.StatusUnauthorized})
	}
	accessToken , err := loginController.LoginUsecase.CreateAccessToken(&user,loginController.Env.AccessTokenSecretKey,loginController.Env.AccessTokenExpiryHour)
	if err != nil{
		c.JSON(http.StatusInternalServerError,domain.ErrorResponse{Message: err.Error(), StatusCode: http.StatusInternalServerError})
	}
	refreshToken ,err := loginController.LoginUsecase.CreateRefreshToken(&user, loginController.Env.RefreshTokenSecretKey, loginController.Env.RefreshTokenExpiryHour)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error(), StatusCode: http.StatusInternalServerError})
	
	}
	loginResponse := domain.LoginResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, loginResponse)

}