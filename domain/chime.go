package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionChime = "chimes"
)

type Chime struct {
	ID           bson.ObjectID `bson:"_id"`
	ChimeTitle   string        `bson:"chime_title"`
	ChimeContent string        `bson:"chime_content"`
	CreatedAt    bson.DateTime `bson:"created_at" swaggertype:"primitive,string"`
	Author       bson.ObjectID `bson:"author"`
	IsPrivate    bool          `bson:"is_private"`
}
type ChimeCreateRequest struct {
	ChimeTitle   string `json:"chime_title"`
	ChimeContent string `json:"chime_content"`
	IsPrivate    bool   `json:"is_private"`
}

type ChimeResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Chime      *Chime `json:"chime"`
}
type ChimeListResponse struct {
	Message    string  `json:"message"`
	StatusCode int     `json:"status_code"`
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
