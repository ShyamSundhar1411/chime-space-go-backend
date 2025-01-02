package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionChime = "chime"
)

type Chime struct {
	ID           primitive.ObjectID `bson:"_id"`
	ChimeTitle   string             `bson:"chime_title"`
	ChimeContent string             `bson:"chime_content"`
	CreatedAt    string             `bson:"created_at"`
}

type ChimeRepository interface {
	Create(c context.Context, chime *Chime) error
	Fetch(c context.Context) ([]Chime, error)
	GetById(c context.Context, id string) (Chime, error)
}
