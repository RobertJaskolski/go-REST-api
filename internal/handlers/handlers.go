package handlers

import "github.com/RobertJaskolski/go-REST-api/internal/repositories"

type Handlers struct {
	UserHandler *UserHandler
}

func NewHandlers(repositories *repositories.Repositories) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(*repositories.UserRepository),
	}
}
