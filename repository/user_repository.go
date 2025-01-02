package repository

import (
	"context"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
	"github.com/ShyamSundhar1411/chime-space-go-backend/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewUserRepository(db mongo.Database,collection string) models.UserRepository{
	return &userRepository{
		database: db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context,user *models.User)error{
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(c,user)
	return err
}

func (ur *userRepository)Fetch(c context.Context)([]models.User,error){
	collection := ur.database.Collection(ur.collection)
	opts := options.Find().SetProjection(bson.D{{Key:"password",Value:0}})
	cursor,err := collection.Find(c, bson.D{}, opts)
	if err != nil{
		return nil, err
	}
	var users []models.User
	err = cursor.All(c, &users)
	if users == nil{
		return []models.User{}, nil
	}
	return users,err
}

func (ur *userRepository)GetById(c context.Context, id string)(models.User, error){
	collection := ur.database.Collection(ur.collection)
	var user models.User
	err := collection.FindOne(c, bson.D{{Key:"_id", Value:id}}).Decode(&user)
	return user, err
}
func (ur *userRepository)GetByUsername(c context.Context, username string)(models.User,error){
	collection := ur.database.Collection(ur.collection)
	var user models.User
	err := collection.FindOne(c, bson.D{{Key:"username",Value: username}}).Decode(&user)
	return user,err
}