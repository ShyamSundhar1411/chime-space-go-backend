package repository

import (
	"context"
	"fmt"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
	"github.com/ShyamSundhar1411/chime-space-go-backend/mongo"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *models.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(c, user)
	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]models.User, error) {
	collection := ur.database.Collection(ur.collection)
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = cursor.All(c, &users)
	if users == nil {
		return []models.User{}, nil
	}
	return users, err
}

func (ur *userRepository) GetById(c context.Context, id string) (*models.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user models.User
	primitiveUserId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: primitiveUserId}}).Decode(&user)
	return &user, err
}
func (ur *userRepository) GetByUsername(c context.Context, username string) (models.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user models.User
	err := collection.FindOne(c, bson.D{{Key: "username", Value: username}}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetMyProfile(c context.Context) (*models.User, error) {
	collection := ur.database.Collection(ur.collection)
	userId, ok := c.Value(utils.UserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("Unable to Fetch User Profile")
	}
	primitiveUserId, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: primitiveUserId}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
