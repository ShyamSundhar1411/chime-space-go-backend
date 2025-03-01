package usecase

import (
	"context"
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

func NewUserUseCase(userRepository domain.UserRepository,timeout time.Duration) domain.UserUsecase{
	return &userUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (cu *userUseCase) GetMyProfile(c context.Context)(*models.User, error){
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.userRepository.GetMyProfile(ctx)
}