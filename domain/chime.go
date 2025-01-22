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
type ChimeCreateRequest struct {
	ChimeTitle   string `json:"chimeTitle" form:"chimeTitle" binding:"required"`
	ChimeContent string `json:"chimeContent" form:"chimeContent" binding:"required"`
	IsPrivate    bool   `json:"isPrivate" form:"isPrivate" binding:"required"`
}

type ChimeResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Chime      *Chime `json:"chime"`
}
type ChimeListResponse struct {
	Message    string  `json:"message"`
	StatusCode int     `json:"statusCode"`
	Chimes     []Chime `json:"chimes"`
}

type ChimeRepository interface {
	CreateChime(c context.Context, chime *Chime) (*Chime, error)
	Fetch(c context.Context) ([]Chime, error)
	GetById(c context.Context, id string) (Chime, error)
	GetChimeFromUserId(c context.Context) ([]Chime, error)
}

type ChimeUsecase interface {
	CreateChime(c context.Context, chime ChimeCreateRequest) (*Chime, error)
	Fetch(c context.Context) ([]Chime, error)
	GetById(c context.Context, id string) (Chime, error)
	FetchChimeFromUser(c context.Context) ([]Chime, error)
}
