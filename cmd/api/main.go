package main

import (
	"github.com/RobertJaskolski/go-REST-api/config"
	"github.com/RobertJaskolski/go-REST-api/internal/api"
	"github.com/RobertJaskolski/go-REST-api/internal/handlers"
	"github.com/RobertJaskolski/go-REST-api/internal/repositories"
	"github.com/RobertJaskolski/go-REST-api/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	// LOAD ENVIRONMENT VARIABLES
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cfg := config.NewConfig()
	err = cfg.Load()
	if err != nil {
		panic(err)
	}

	// CREATE NEW DATABASE CONNECTION
	dbPool, err := db.NewPostgresConnection(cfg)
	if err != nil {
		panic(err)
	}

	defer dbPool.Close()

	// INITIALIZE REPOSITORIES
	r := repositories.NewRepositories(dbPool)

	// INITIALIZE HANDLERS
	h := handlers.NewHandlers(cfg, r)

	// CREATE NEW SERVER
	server := api.NewServer(cfg, dbPool)

	// SETUP ROUTES
	server.SetupRoutes(h)

	// MIDDLEWARE
	server.SetupValidator()

	// START SERVER
	err = server.RunAndListen()
	if err != nil {
		panic(err)
	}
}
