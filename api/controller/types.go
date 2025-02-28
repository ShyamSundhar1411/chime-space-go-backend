package controller

import (
	"github.com/ShyamSundhar1411/chime-space-go-backend/bootstrap"
	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
)

type ChimeController struct {
	ChimeUsecase domain.ChimeUsecase
}

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}
type SignUpController struct {
	SignUpUsecase domain.SignUpUsecase
	Env           *bootstrap.Env
}

type UserController struct{
	UserUseCase domain.UserUsecase
	Env *bootstrap.Env
}

type TokenController struct{
	TokenUseCase domain.TokenUsecase
	Env *bootstrap.Env
}