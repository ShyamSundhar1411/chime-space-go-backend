package repository

import (
	"context"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
	"github.com/ShyamSundhar1411/chime-space-go-backend/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewProfileRepository(db mongo.Database, collection string) domain.ProfileRepository {
	return &profileRepository{
		database:   db,
		collection: collection,
	}
}

func (pr *profileRepository) GetProfileByUser(ctx context.Context, user *models.User) (*models.Profile, error) {
	collection := pr.database.Collection(pr.collection)
	var profile models.Profile
	err := collection.FindOne(ctx, bson.M{"user": user.ID}).Decode(&profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
