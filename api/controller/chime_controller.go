package controller

import (
	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)
// FetchAllChimes handles the endpoint for fetching all chimes
// @Summary Get all Chimes
// @Description Fetch all Chimes from the database
// @Tags Chimes
// @Accept json
// @Produce json
// @Success 200 {array} domain.Chime
// @Failure 500 {object} domain.ErrorResponse
// @Router /chimes [get]
func (chimeController *ChimeController) FetchAllChimes(c echo.Context) error {
	var chimes []domain.Chime
	ctx := utils.ExtractContext(c)
	chimes, err := chimeController.ChimeUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	return c.JSON(http.StatusOK, chimes)
}
