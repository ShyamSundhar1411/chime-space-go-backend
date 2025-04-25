package usecase

import (
	"context"
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
)

func NewUserUseCase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (cu *userUseCase) GetMyProfile(c context.Context) (*domain.UserProfile, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	user, err := cu.userRepository.GetMyProfile(ctx)
	if err != nil {
		return nil, err
	}
	userProfile := &domain.UserProfile{
		User: user,
	}
	return userProfile, nil
}
