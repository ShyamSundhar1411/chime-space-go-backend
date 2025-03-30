package usecase

import (
	"context"
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
)

func NewTokenUseCase(userRepository domain.UserRepository,timeout time.Duration) domain.TokenUsecase{
	return &tokenUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func(tokenUseCase *tokenUseCase) GenerateAccessToken(c context.Context, user *models.User, secret string, expiry int)(string,error){
	t, err := utils.CreateAccessToken(user, secret, expiry)
	if err != nil{
		return "", err
	}
	return t,nil	
}

func(tokenUseCase *tokenUseCase) GenerateRefreshToken(c context.Context, user *models.User, secret string, expiry int)(string, error){
	t, err := utils.CreateRefreshToken(user, secret, expiry)
	if err != nil{
		return "", err
	}
	return t,nil
}
func(tokenUseCase *tokenUseCase)ValidateRefreshToken(c context.Context, refreshToken string, secret string)(*models.User, error){
	claims, err := utils.VerifyRefreshToken(refreshToken, secret)
	if err != nil{
		return nil, err
	}
	user, err := tokenUseCase.userRepository.GetById(c, claims.ID)
	if err != nil{
		return nil, err
	}
	return user,nil
}