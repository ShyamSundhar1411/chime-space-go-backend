package controller

import (
	"net/http"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"github.com/labstack/echo/v4"
)

// FetchAllChimes handles the endpoint for fetching all chimes
//
//	@Summary		Get all Chimes
//	@Description	Fetch all Chimes from the database
//	@Tags			Chimes
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		domain.ChimeListResponse
//	@Failure		500	{object}	domain.ChimeListResponse
//	@Router			/chimes/ [get]
func (chimeController *ChimeController) FetchAllChimes(c echo.Context) error {
	var chimes []domain.Chime
	ctx := utils.ExtractContext(c)
	chimes, err := chimeController.ChimeUsecase.Fetch(ctx)
	if err != nil {
		response := domain.ChimeListResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			Chimes:     nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := domain.ChimeListResponse{
		Message:    "Chimes fetched successfully",
		StatusCode: http.StatusOK,
		Chimes:     chimes,
	}
	return c.JSON(http.StatusOK, response)
}

// FetchAllChimes authored by the logged in user
//
//	@Summary		Get all Chimes of logged in user
//	@Description	Fetch all Chimes from logged in user from the database
//	@Tags			Chimes
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		domain.Chime
//	@Failure		500	{object}	domain.ChimeListResponse
//	@Router			/chimes/user/ [get]
//	@Security		BearerAuth
func (chimeController *ChimeController) FetchChimeFromUser(c echo.Context) error {
	var chimes []domain.Chime
	ctx := utils.ExtractContext(c)
	chimes, err := chimeController.ChimeUsecase.FetchChimeFromUser(ctx)
	if err != nil {
		response := domain.ChimeListResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Chimes:     nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response := domain.ChimeListResponse{
		Message:    "Chimes fetched successfully",
		StatusCode: http.StatusOK,
		Chimes:     chimes,
	}
	return c.JSON(http.StatusOK, response)
}

// Creates a Chime
//
//	@Summary		Creates a new Chime
//	@Description	Create a new chime by providing a request body with necessary details.
//	@Tags			Chimes
//	@Accept			json
//	@Produce		json
//	@Param			request	body		domain.ChimeCreateRequest	true	"Chime Create Request"
//	@Success		201		{object}	domain.Chime
//	@Failure		400		{object}	domain.ChimeResponse
//	@Router			/chimes/ [post]
//	@Security		BearerAuth
func (ChimeController *ChimeController) CreateChime(c echo.Context) error {
	var chime *domain.Chime
	var request domain.ChimeCreateRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	ctx := utils.ExtractContext(c)
	chime, err = ChimeController.ChimeUsecase.CreateChime(ctx, request)
	if err != nil {
		chimeResponse := domain.ChimeResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Chime: chime}
		return c.JSON(http.StatusBadRequest, chimeResponse)
	}
	chimeResponse := domain.ChimeResponse{Message: "Chime created successfully", StatusCode: http.StatusCreated, Chime: chime}
	return c.JSON(http.StatusCreated, chimeResponse)
}
