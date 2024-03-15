package config

import (
	"flag"
	"os"
)

type Config struct {
	API struct {
		Port string
	}
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) Load() error {
	// API Config
	flag.StringVar(&cfg.API.Port, "port", os.Getenv("API_PORT"), "API Server Port")

	// DB Config
	flag.StringVar(&cfg.DB.Host, "db_host", os.Getenv("DB_HOST"), "Database Host")
	flag.StringVar(&cfg.DB.Port, "db_port", os.Getenv("DB_PORT"), "Database Port")
	flag.StringVar(&cfg.DB.User, "db_user", os.Getenv("DB_USER"), "Database User")
	flag.StringVar(&cfg.DB.Password, "db_password", os.Getenv("DB_PASSWORD"), "Database Password")
	flag.StringVar(&cfg.DB.Database, "db_database", os.Getenv("DB_DATABASE"), "Database Name")

	// Parse the flags
	flag.Parse()

	return nil
}
