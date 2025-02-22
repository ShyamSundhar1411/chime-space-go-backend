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
	var chimes []domain.ChimeWithAuthor
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
	var chimes []domain.ChimeWithAuthor
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
//	@Param			request	body		domain.ChimeCreateOrUpdateRequest	true	"Chime Create Request"
//	@Success		201		{object}	domain.Chime
//	@Failure		400		{object}	domain.ChimeResponse
//	@Router			/chimes/ [post]
//	@Security		BearerAuth
func (ChimeController *ChimeController) CreateChime(c echo.Context) error {
	var chime *domain.ChimeWithAuthor
	var request domain.ChimeCreateOrUpdateRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
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

// UpdateChime updates an existing Chime
//
//	@Summary		Update an existing Chime
//	@Description	Update an existing chime by providing the chime ID and updated details
//	@Tags			Chimes
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string							true	"Chime ID"
//	@Param			request	body		domain.ChimeCreateOrUpdateRequest	true	"Chime Update Request"
//	@Success		200		{object}	domain.Chime
//	@Failure		400		{object}	domain.ChimeResponse
//	@Router			/chimes/{id} [put]
//	@Security		BearerAuth
func (ChimeController *ChimeController) UpdateChime(c echo.Context) error {
	var chime *domain.ChimeWithAuthor
	var request domain.ChimeCreateOrUpdateRequest
	id := c.Param("id")
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	ctx := utils.ExtractContext(c)
	chime, err = ChimeController.ChimeUsecase.UpdateChime(ctx, request, id)
	if err != nil {
		chimeResponse := domain.ChimeResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Chime: chime}
		return c.JSON(http.StatusBadRequest, chimeResponse)
	}
	chimeResponse := domain.ChimeResponse{Message: "Chime updated successfully", StatusCode: http.StatusOK, Chime: chime}
	return c.JSON(http.StatusOK, chimeResponse)
}
// DeleteChime Deletes an existing Chime
//
//	@Summary		Delete an existing Chime
//	@Description	Delete an existing chime by providing the chime ID
//	@Tags			Chimes
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string							true	"Chime ID"
//	@Success		200		{object}	domain.BaseResponse
//	@Failure		404		{object}	domain.BaseResponse
//	@Router			/chimes/{id} [delete]
//	@Security		BearerAuth
func (ChimeController *ChimeController) DeleteChime(c echo.Context) error {
	id := c.Param("id")
	ctx := utils.ExtractContext(c)
	err := ChimeController.ChimeUsecase.DeleteChime(ctx,id)
	if err != nil{
		return c.JSON(http.StatusNotFound, domain.BaseResponse{Message: "Unable to Process Request", StatusCode: http.StatusNotFound})
	}
	return c.JSON(http.StatusOK, domain.BaseResponse{Message: "Chime deleted successfully", StatusCode: http.StatusOK})
}