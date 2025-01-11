package usecase

import (
	"context"
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
)

func NewChimeUseCase(chimeRepository domain.ChimeRepository, timeout time.Duration) domain.ChimeUsecase {
	return &chimeUsecase{
		chimeRepository: chimeRepository,
		contextTimeout:  timeout,
	}
}

func (cu *chimeUsecase) Create(c context.Context, chime *domain.Chime) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chimeRepository.Create(ctx, chime)
}

func (cu *chimeUsecase) Fetch(c context.Context) ([]domain.Chime, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chimeRepository.Fetch(ctx)
}

func (cu *chimeUsecase) GetById(c context.Context, id string) (domain.Chime, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chimeRepository.GetById(ctx, id)
}
