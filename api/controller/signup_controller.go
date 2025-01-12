package controller

import (
	"net/http"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func (signUpController *SignUpController) SignUp(c echo.Context) error {
	var request domain.SignUpRequest
	err := c.Bind(&request)
	ctx := utils.ExtractContext(c)
	if err != nil{
		return c.JSON(http.StatusConflict,domain.ErrorResponse{Message: "Invalid Request",StatusCode: http.StatusConflict})
	}
	_, err = signUpController.SignUpUsecase.GetUserByUsername(ctx, request.UserName)
	if err == nil{
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid Request", StatusCode: http.StatusBadRequest})
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Internal Server Error", StatusCode: http.StatusInternalServerError})
	}
	request.Password = string(encryptedPassword)
	user := domain.User{
		ID:       bson.NewObjectID(),
		UserName: request.UserName,
		Email: request.Email,
		PenName: request.PenName,
		Password: request.Password,
	}
	err = signUpController.SignUpUsecase.Create(ctx,&user)
	if err != nil{
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Bad Request", StatusCode: http.StatusBadRequest})
	}
	accessToken,err := signUpController.SignUpUsecase.CreateAccessToken(&user,signUpController.Env.AccessTokenSecretKey,signUpController.Env.AccessTokenExpiryHour)
	if err != nil{
		return c.JSON(http.StatusInternalServerError,domain.ErrorResponse{Message: err.Error(), StatusCode: http.StatusInternalServerError})
	}
	refreshToken ,err := signUpController.SignUpUsecase.CreateRefreshToken(&user, signUpController.Env.RefreshTokenSecretKey, signUpController.Env.RefreshTokenExpiryHour)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error(), StatusCode: http.StatusInternalServerError})
	
	}
	signupResponse := domain.SignUpResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(http.StatusCreated, signupResponse)
}