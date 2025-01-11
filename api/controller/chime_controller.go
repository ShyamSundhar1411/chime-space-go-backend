package controller

import (
	"net/http"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/labstack/echo/v4"
)

func (chimeController *ChimeController) FetchAllChimes(c echo.Context) error {
	var chimes []domain.Chime
	ctx := utils.ExtractContext(c)
	chimes, err := chimeController.ChimeUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	return c.JSON(http.StatusOK, chimes)
}
