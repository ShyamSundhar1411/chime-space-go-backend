package usecase

import (
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
)

type chimeUsecase struct {
	chimeRepository domain.ChimeRepository
	contextTimeout  time.Duration
}

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}
