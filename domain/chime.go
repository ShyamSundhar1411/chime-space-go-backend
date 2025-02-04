package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionChime = "chimes"
)

type Chime struct {
	ID           bson.ObjectID `bson:"_id" json:"id"`
	ChimeTitle   string        `bson:"chime_title" json:"chimeTitle"`
	ChimeContent string        `bson:"chime_content" json:"chimeContent"`
	CreatedAt    bson.DateTime `bson:"created_at" json:"createdAt" swaggertype:"primitive,string"`
	Author       bson.ObjectID `bson:"author" json:"author"`
	IsPrivate    bool          `bson:"is_private" json:"isPrivate"`
}

type ChimeWithAuthor struct {
	Chime  `json:",inline" bson:",inline"`
	Author User `json:"author"`
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
	CreateChime(c context.Context, chime *Chime) (*ChimeWithAuthor, error)
	Fetch(c context.Context) ([]ChimeWithAuthor, error)
	GetById(c context.Context, id string) (*ChimeWithAuthor, error)
	GetChimeFromUserId(c context.Context) ([]ChimeWithAuthor, error)
	UpdateChime(c context.Context, chimeData ChimeCreateOrUpdateRequest, id string) (*ChimeWithAuthor, error)
}

type ChimeUsecase interface {
	CreateChime(c context.Context, chime ChimeCreateOrUpdateRequest) (*ChimeWithAuthor, error)
	Fetch(c context.Context) ([]ChimeWithAuthor, error)
	GetById(c context.Context, id string) (*ChimeWithAuthor, error)
	FetchChimeFromUser(c context.Context) ([]ChimeWithAuthor, error)
	UpdateChime(c context.Context, chime ChimeCreateOrUpdateRequest, id string) (*ChimeWithAuthor, error)
}
