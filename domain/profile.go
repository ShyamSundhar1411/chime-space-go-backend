package domain

import (
	"context"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

const (
	CollectionProfile = "profiles"
)

type ProfileRepository interface {
	GetProfileByUser(c context.Context, user *models.User) (*models.Profile, error)
}
