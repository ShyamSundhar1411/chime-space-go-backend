package usecase

import (
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

type chimeUsecase struct {
	chimeRepository models.ChimeRepository
	contextTimeout time.Duration
}