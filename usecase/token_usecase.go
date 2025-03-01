package usecase

import (
	"context"
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
)

func NewTokenUseCase(userRepository domain.UserRepository,timeout time.Duration) domain.TokenUsecase{
	return &tokenUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func(tokenUseCase *tokenUseCase) RefreshToken(c context.Context, refreshToken string)(string,error){
	return "",nil
}

func(tokenUseCase *tokenUseCase)ValidateRefreshToken(c context.Context, refreshToken string, secret string)(bool, error){
	claims, err := utils.VerifyRefreshToken(refreshToken, secret)
	if err != nil{
		return false, err
	}
	_, err = tokenUseCase.userRepository.GetById(c, claims.ID)
	if err != nil{
		return false, err
	}
	return true,nil
}