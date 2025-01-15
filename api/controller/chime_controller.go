package controller

import (
	"net/http"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// FetchAllChimes handles the endpoint for fetching all chimes
//
//	@Summary		Get all Chimes
//	@Description	Fetch all Chimes from the database
//	@Tags			Chimes
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		domain.Chime
//	@Failure		500	{object}	domain.ErrorResponse
//	@Router			/chimes/ [get]
func (chimeController *ChimeController) FetchAllChimes(c echo.Context) error {
	var chimes []domain.Chime
	ctx := utils.ExtractContext(c)
	chimes, err := chimeController.ChimeUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	return c.JSON(http.StatusOK, chimes)
}
// FetchAllChimes authored by the logged in user
//
//	@Summary		Get all Chimes of logged in user
//	@Description	Fetch all Chimes from logged in user from the database
//	@Tags			Chimes
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		domain.Chime
//	@Failure		500	{object}	domain.ErrorResponse
//	@Router			/chimes/user/ [get]
//	@Security		BearerAuth
func (chimeController *ChimeController) FetchChimeFromUser(c echo.Context) error{
	user := c.Get("user").(*jwt.Token)
	return c.JSON(http.StatusOK,user)
}