package usecase

import (
	"context"
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
)

func NewLoginUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

func (lu *loginUsecase) GetUserByUserName(c context.Context, userName string) (models.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByUsername(ctx, userName)
}
func (lu *loginUsecase) CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}
func (lu *loginUsecase) CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
