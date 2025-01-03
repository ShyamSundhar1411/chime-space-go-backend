package usecase

import (
	"context"
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

func NewChimeUseCase(chimeRepository models.ChimeRepository, timeout time.Duration) models.ChimeUsecase {
	return &chimeUsecase{
		chimeRepository: chimeRepository,
		contextTimeout:  timeout,
	}
}

func (cu *chimeUsecase) Create(c context.Context, chime *models.Chime) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chimeRepository.Create(ctx, chime)
}

func (cu *chimeUsecase) Fetch(c context.Context) ([]models.Chime, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chimeRepository.Fetch(ctx)
}

func (cu *chimeUsecase) GetById(c context.Context, id string) (models.Chime, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chimeRepository.GetById(ctx, id)
}
