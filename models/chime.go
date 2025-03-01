package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Chime struct {
	ID           bson.ObjectID `bson:"_id" json:"id"`
	ChimeTitle   string        `bson:"chime_title" json:"chimeTitle"`
	ChimeContent string        `bson:"chime_content" json:"chimeContent"`
	CreatedAt    bson.DateTime `bson:"created_at" json:"createdAt" swaggertype:"primitive,string"`
	Author       bson.ObjectID `bson:"author" json:"author"`
	IsPrivate    bool          `bson:"is_private" json:"isPrivate"`
}