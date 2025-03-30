package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
	"github.com/ShyamSundhar1411/chime-space-go-backend/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func buildChimePipeline(matchFilter bson.M) []bson.M {
	return []bson.M{
		{"$match": matchFilter},
		{"$lookup": bson.M{
			"from":         "users",
			"localField":   "author",
			"foreignField": "_id",
			"as":           "author",
		}},
		{"$sort": bson.M{"created_at": -1}},
		{"$unwind": "$author"},
		{
			"$project": bson.M{
				"_id":             1,
				"chime_title":     1,
				"chime_content":   1,
				"created_at":      1,
				"is_private":      1,
				"author._id":      1,
				"author.username": 1,
				"author.penname":  1,
				"author.email":    1,
			},
		},
	}
}
func NewChimeRepository(db mongo.Database, collection string) domain.ChimeRepository {
	return &chimeRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *chimeRepository) CreateChime(c context.Context, chime *models.Chime) (*domain.ChimeWithAuthor, error) {
	collection := cr.database.Collection(cr.collection)

	if chime.ID.IsZero() {
		chime.ID = bson.NewObjectID()
	}

	_, err := collection.InsertOne(c, chime)
	if err != nil {
		return nil, err
	}

	pipeline := buildChimePipeline(bson.M{"_id": chime.ID})

	cursor, err := collection.Aggregate(c, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var result domain.ChimeWithAuthor
	if cursor.Next(c) {
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("no chime found with the given ID")
	}

	return &result, nil
}

func (cr *chimeRepository) Fetch(c context.Context) ([]domain.ChimeWithAuthor, error) {
	collection := cr.database.Collection(cr.collection)
	pipeline := buildChimePipeline(bson.M{"is_private": false})
	cursor, err := collection.Aggregate(c, pipeline)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var chimes []domain.ChimeWithAuthor
	err = cursor.All(c, &chimes)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if chimes == nil {
		return []domain.ChimeWithAuthor{}, nil
	}
	return chimes, err

}

func (cr *chimeRepository) GetById(c context.Context, id string) (*domain.ChimeWithAuthor, error) {
	collection := cr.database.Collection(cr.collection)
	var chime domain.ChimeWithAuthor
	primitiveID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	pipeline := buildChimePipeline(bson.M{"_id": primitiveID})
	cursor, err := collection.Aggregate(c, pipeline)
	if err != nil {
		return nil,err
	}
	err = cursor.Decode(&chime)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &chime, err
}

func (cr *chimeRepository) GetChimeFromUserId(c context.Context) ([]domain.ChimeWithAuthor, error) {
	collection := cr.database.Collection(cr.collection)
	userId, ok := c.Value("userId").(string)
	if !ok || userId == "" {
		return nil, fmt.Errorf("Invalid user id ")
	}
	primitiveUserId, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	pipeline := buildChimePipeline(bson.M{"author": primitiveUserId})
	cursor, err := collection.Aggregate(c, pipeline)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var chimes []domain.ChimeWithAuthor
	err = cursor.All(c, &chimes)
	if err != nil {
		return nil, err
	}
	if chimes == nil {
		return []domain.ChimeWithAuthor{}, nil
	}
	return chimes, err
}

func (cr *chimeRepository) UpdateChime(c context.Context, chimeData domain.ChimeCreateOrUpdateRequest, id string) (*domain.ChimeWithAuthor, error) {
	collection := cr.database.Collection(cr.collection)
	userId, ok := c.Value("userId").(string)
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
	filter := bson.D{{Key: "_id", Value: prmitiveChimeId}, {Key: "author", Value: primitiveUserId}}

	update := bson.M{
		"$set": bson.M{
			"chime_title":   chimeData.ChimeTitle,
			"chime_content": chimeData.ChimeContent,
			"is_private":    chimeData.IsPrivate,
		},
	}
	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("Chime not found")
	}

	var updatedChime domain.ChimeWithAuthor
	pipeline := buildChimePipeline(bson.M{"_id": prmitiveChimeId})
	cursor, err := collection.Aggregate(c, pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.Decode(&updatedChime)
	if err != nil {
		return nil, err
	}
	return &updatedChime, nil
}

func(cr *chimeRepository) DeleteChime(c context.Context,id string)(error){
	collection := cr.database.Collection(cr.collection)
	userId, ok := c.Value("userId").(string)
	if !ok || userId == "" {
		return fmt.Errorf("Invalid user id ")
	}
	primitiveUserId, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	primitiveChimeId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": primitiveChimeId, "author": primitiveUserId})
	if err != nil {
		return err
	}
	return nil
}