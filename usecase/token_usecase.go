package usecase

import (
	"context"
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
)

func NewTokenUseCase(timeout time.Duration) domain.TokenUsecase{
	return &tokenUseCase{
		contextTimeout: timeout,
	}
}

func(tokenUseCase *tokenUseCase) RefreshToken(c context.Context, refreshToken string)(string,error){
	return "",nil
}