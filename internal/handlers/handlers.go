package handlers

import (
	"github.com/RobertJaskolski/go-REST-api/config"
	"github.com/RobertJaskolski/go-REST-api/internal/repositories"
)

type Handlers struct {
	UserHandler *UserHandler
	AuthHandler *AuthHandler
}

func NewHandlers(cfg *config.Config, repositories *repositories.Repositories) *Handlers {
	return &Handlers{
		AuthHandler: NewAuthHandler(cfg, *repositories.UserRepository),
		UserHandler: NewUserHandler(*repositories.UserRepository),
	}
}
