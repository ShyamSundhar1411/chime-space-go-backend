package repository

import "github.com/ShyamSundhar1411/chime-space-go-backend/mongo"

type userRepository struct {
	database   mongo.Database
	collection string
}

type chimeRepository struct {
	database   mongo.Database
	collection string
}

type profileRepository struct {
	database   mongo.Database
	collection string
}
