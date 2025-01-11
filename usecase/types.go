package usecase

import (
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/models"
)

type chimeUsecase struct {
	chimeRepository models.ChimeRepository
	contextTimeout  time.Duration
}

type signupUsecase struct {
	userRepository models.UserRepository
	contextTimeout time.Duration
}

type loginUsecase struct{
	userRepository models.UserRepository
	contextTimeout time.Duration
}