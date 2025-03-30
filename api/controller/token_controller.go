package controller

import (
	"fmt"
	"net/http"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"github.com/labstack/echo/v4"
)

// @Summary      Generate a new access token using a refresh token
// @Description  Allows users to refresh their access token using a valid refresh token.
// @Tags         Token
// @Accept       json
// @Produce      json
// @Param        refreshRequest  body  domain.TokenRefreshRequest  true  "Refresh Token Request Payload"
// @Success      200  {object}  domain.TokenRefreshResponse  "Returns a new access token"
// @Success      400  {object}  domain.BaseResponse  "Bad Payload"
// @Success      401  {object}  domain.BaseResponse  "Unauthorized"
// @Router       /token/refresh/ [post]
func (tokenController *TokenController) Refresh(c echo.Context) error {
	var request domain.TokenRefreshRequest
	err := c.Bind(&request)
	if err != nil{
		return c.JSON(http.StatusBadRequest, domain.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message: "Invalid payload",
		})
	}
	ctx := utils.ExtractContext(c)
	user, err := tokenController.TokenUseCase.ValidateRefreshToken(ctx,request.RefreshToken, tokenController.Env.RefreshTokenSecretKey)
	fmt.Printf("Error: %v",err)
	if err != nil{
		return c.JSON(http.StatusUnauthorized, domain.BaseResponse{
			StatusCode: http.StatusUnauthorized,
			Message: "Invalid refresh token",
		})
	}
	accessToken, err := tokenController.TokenUseCase.GenerateAccessToken(ctx, user, tokenController.Env.AccessTokenSecretKey,tokenController.Env.AccessTokenExpiryHour)
	if err != nil{
		return c.JSON(http.StatusUnauthorized, domain.BaseResponse{
			StatusCode: http.StatusUnauthorized,
			Message: "Failed to generate token",
		})
	}
	refreshToken, err := tokenController.TokenUseCase.GenerateRefreshToken(ctx, user, tokenController.Env.RefreshTokenSecretKey, tokenController.Env.RefreshTokenExpiryHour)
	if err != nil{
		return c.JSON(http.StatusUnauthorized, domain.BaseResponse{
			StatusCode: http.StatusUnauthorized,
			Message: "Failed to generate token",
		})
	}
	
	tokenResponse := domain.TokenRefreshResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		StatusCode: http.StatusCreated,
		Message: "Token refreshed successfully",
	}
	return c.JSON(
		http.StatusOK,
		tokenResponse,
	)
}
