package usecase

import (
	"context"
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/entities"
	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
	"github.com/ShyamSundhar1411/chime-space-go-backend/utils"
)



func NewSignUpUsecase(userRepository models.UserRepository, timeout time.Duration) entities.SignUpUsecase{
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context,user *models.User) error {
	ctx,cancel := context.WithTimeout(c,su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByUsername(c context.Context,username string) (models.User, error) {
	ctx,cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByUsername(ctx,username)
}

func (su *signupUsecase) CreateAccessToken(user *models.User,secret string,expiry int)(accessToken string,err error){
	return utils.CreateAccessToken(user,secret,expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *models.User, secret string, expiry int)(refreshToken string, err error){
	return utils.CreateRefreshToken(user, secret, expiry)
}
	