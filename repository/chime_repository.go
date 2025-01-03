package repository

import (
	"context"
	"log"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
	"github.com/ShyamSundhar1411/chime-space-go-backend/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewChimeRepository(db mongo.Database, collection string) models.ChimeRepository {
	return &chimeRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *chimeRepository) Create(c context.Context, chime *models.Chime) error {
	collection := cr.database.Collection(cr.collection)
	_, err := collection.InsertOne(c, chime)
	return err

}

func (cr *chimeRepository) Fetch(c context.Context) ([]models.Chime, error) {
	collection := cr.database.Collection(cr.collection)
	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var chimes []models.Chime
	err = cursor.All(c, &chimes)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if chimes == nil {
		return []models.Chime{}, nil
	}
	return chimes, err

}

func (cr *chimeRepository) GetById(c context.Context, id string) (models.Chime, error) {
	collection := cr.database.Collection(cr.collection)
	var chime models.Chime
	err := collection.FindOne(c, bson.D{{Key: "_id", Value: id}}).Decode(&chime)
	return chime, err
}
