package controller

import (
	"net/http"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"github.com/labstack/echo/v4"
)

func (chimeController *ChimeController) FetchAllChimes(c echo.Context) error{
	var chimes []models.Chime
	ctx := utils.ExtractContext(c)
	chimes,err := chimeController.ChimeUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,utils.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, chimes)
}