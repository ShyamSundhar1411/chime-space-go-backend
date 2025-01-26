package repository

import (
	"context"
	"fmt"
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

func (cr *chimeRepository) CreateChime(c context.Context, chime *domain.Chime) (*domain.Chime, error) {
	collection := cr.database.Collection(cr.collection)
	_, err := collection.InsertOne(c, chime)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: chime.ID}}).Decode(&chime)
	return chime, err

}

func (cr *chimeRepository) Fetch(c context.Context) ([]domain.Chime, error) {
	collection := cr.database.Collection(cr.collection)
	cursor, err := collection.Find(c, bson.D{{Key: "is_private", Value: false}})
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

func (cr *chimeRepository) GetChimeFromUserId(c context.Context) ([]domain.Chime, error) {
	collection := cr.database.Collection(cr.collection)
	userId, ok := c.Value("userId").(string)
	if !ok || userId == "" {
		return nil, fmt.Errorf("Invalid user id ")
	}
	primitiveUserId, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	cursor, err := collection.Find(c, bson.D{{Key: "author", Value: primitiveUserId}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var chimes []domain.Chime
	err = cursor.All(c, &chimes)
	if err != nil {
		return nil, err
	}
	if chimes == nil {
		return []domain.Chime{}, nil
	}
	return chimes, err
}

func(cr *chimeRepository) UpdateChime(c context.Context, chimeData domain.ChimeCreateOrUpdateRequest, id string)(*domain.Chime, error) {
	collection := cr.database.Collection(cr.collection)
	userId,ok := c.Value("userId").(string)
	if !ok || userId == "" {
		return nil, fmt.Errorf("Invalid user id ")
	}
	primitiveUserId, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	prmitiveChimeId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: prmitiveChimeId},{Key:"author",Value:primitiveUserId}}
	
	update := bson.M{
		"$set": bson.M{
			"chime_title":   chimeData.ChimeTitle,
			"chime_content": chimeData.ChimeContent,
			"is_private":    chimeData.IsPrivate,
		},
	}
	result,err := collection.UpdateOne(c, filter , update)	
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("Chime not found")
	}
	
	var updatedChime domain.Chime
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: prmitiveChimeId}}).Decode(&updatedChime)
	if err != nil {
		return nil, err
	}
	return &updatedChime, nil
}