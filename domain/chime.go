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
	CreatedAt    string        `bson:"created_at"`
	Author 		 bson.ObjectID  `bson:"author"`
	IsPrivate	 bool			`bson:"is_private"`
}

type ChimeRepository interface {
	Create(c context.Context, chime *Chime) error
	Fetch(c context.Context) ([]Chime, error)
	GetById(c context.Context, id string) (Chime, error)
	GetChimeFromUserId(c context.Context, id string) ([]Chime, error)
}

type ChimeUsecase interface {
	Create(c context.Context, chime *Chime) error
	Fetch(c context.Context) ([]Chime, error)
	GetById(c context.Context, id string) (Chime, error)
	FetchChimeFromUser(c context.Context, id string) ([]Chime, error)
}
