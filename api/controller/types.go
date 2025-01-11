package controller

import (
	"github.com/ShyamSundhar1411/chime-space-go-backend/bootstrap"
	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
)

type ChimeController struct {
	ChimeUsecase domain.ChimeUsecase
}

type LoginController struct{
	LoginUsecase domain.LoginUsecase
	Env *bootstrap.Env
}
