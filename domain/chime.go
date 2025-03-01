package domain

import (
	"context"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

const (
	CollectionChime = "chimes"
)



type ChimeWithAuthor struct {
	models.Chime  `json:",inline" bson:",inline"`
	Author models.User `json:"author"`
}

type ChimeCreateOrUpdateRequest struct {
	ChimeTitle   string `json:"chimeTitle" form:"chimeTitle" binding:"required"`
	ChimeContent string `json:"chimeContent" form:"chimeContent" binding:"required"`
	IsPrivate    bool   `json:"isPrivate" form:"isPrivate" binding:"required"`
}

type ChimeResponse struct {
	Message    string           `json:"message"`
	StatusCode int              `json:"statusCode"`
	Chime      *ChimeWithAuthor `json:"chime"`
}
type ChimeListResponse struct {
	Message    string            `json:"message"`
	StatusCode int               `json:"statusCode"`
	Chimes     []ChimeWithAuthor `json:"chimes"`
}
type ChimeRepository interface {
	CreateChime(c context.Context, chime *models.Chime) (*ChimeWithAuthor, error)
	Fetch(c context.Context) ([]ChimeWithAuthor, error)
	GetById(c context.Context, id string) (*ChimeWithAuthor, error)
	GetChimeFromUserId(c context.Context) ([]ChimeWithAuthor, error)
	UpdateChime(c context.Context, chimeData ChimeCreateOrUpdateRequest, id string) (*ChimeWithAuthor, error)
	DeleteChime(c context.Context, id string)(error)
}

type ChimeUsecase interface {
	CreateChime(c context.Context, chime ChimeCreateOrUpdateRequest) (*ChimeWithAuthor, error)
	Fetch(c context.Context) ([]ChimeWithAuthor, error)
	GetById(c context.Context, id string) (*ChimeWithAuthor, error)
	FetchChimeFromUser(c context.Context) ([]ChimeWithAuthor, error)
	UpdateChime(c context.Context, chime ChimeCreateOrUpdateRequest, id string) (*ChimeWithAuthor, error)
	DeleteChime(c context.Context, id string)(error)
}
