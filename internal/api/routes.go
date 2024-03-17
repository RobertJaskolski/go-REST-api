package api

import "github.com/RobertJaskolski/go-REST-api/internal/handlers"

func (s *Server) SetupRoutes(h *handlers.Handlers) {
	SetupUserRoutes(s, h)
}

func SetupUserRoutes(s *Server, h *handlers.Handlers) {
	g := s.router.Group("/user")

	// CRUD OPERATIONS FOR USER
	g.POST("/", h.UserHandler.CreateUser)
	g.GET("/", h.UserHandler.GetUsers)
	g.GET("/:id", h.UserHandler.GetUser)
	g.PATCH("/:id", h.UserHandler.UpdateUser)
	g.DELETE("/:id", h.UserHandler.DeleteUser)
}
