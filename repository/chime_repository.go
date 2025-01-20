package repository

import (
	"context"
	"log"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewChimeRepository(db mongo.Database, collection string) domain.ChimeRepository {
	return &chimeRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *chimeRepository) Create(c context.Context, chime *domain.Chime) error {
	collection := cr.database.Collection(cr.collection)
	_, err := collection.InsertOne(c, chime)
	return err

}

func (cr *chimeRepository) Fetch(c context.Context) ([]domain.Chime, error) {
	collection := cr.database.Collection(cr.collection)
	cursor, err := collection.Find(c, bson.D{{Key:"is_private",Value: false}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var chimes []domain.Chime
	err = cursor.All(c, &chimes)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if chimes == nil {
		return []domain.Chime{}, nil
	}
	return chimes, err

}

func (cr *chimeRepository) GetById(c context.Context, id string) (domain.Chime, error) {
	collection := cr.database.Collection(cr.collection)
	var chime domain.Chime
	err := collection.FindOne(c, bson.D{{Key: "_id", Value: id}}).Decode(&chime)
	return chime, err
}

func (cr *chimeRepository) GetChimeFromUserId(c context.Context, userId string)([] domain.Chime, error){
	collection := cr.database.Collection(cr.collection)
	cursor, err := collection.Find(c, bson.D{{Key:"author",Value: userId}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var chimes []domain.Chime
	err = cursor.All(c, &chimes)
	if err!=nil{
		return nil,err
	}
	if chimes == nil {
		return []domain.Chime{}, nil
	}
	return chimes, err
}