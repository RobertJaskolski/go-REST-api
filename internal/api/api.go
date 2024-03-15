package api

import (
	"fmt"
	"github.com/RobertJaskolski/go-REST-api/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"sync"
)

type Server struct {
	port   string
	router *echo.Echo
	db     *pgxpool.Pool
}

var (
	once     sync.Once
	instance *Server
)

func NewServer(cfg *config.Config, db *pgxpool.Pool) *Server {
	once.Do(func() {
		instance = &Server{
			port:   cfg.API.Port,
			router: echo.New(),
			db:     db,
		}
	})

	return instance
}

func (s *Server) RunAndListen() error {
	return s.router.Start(fmt.Sprintf(":%s", s.port))
}
