package db

import (
	"context"
	"fmt"
	"github.com/RobertJaskolski/go-REST-api/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

const (
	maxConn           = 50
	healthCheckPeriod = 3 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
)

func NewPostgresConnection(cfg *config.Config) (*pgxpool.Pool, error) {
	// Create new connection string
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database)

	// Create a new connection pool config
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	// Set up the connection pool
	poolConfig.MaxConns = maxConn
	poolConfig.HealthCheckPeriod = healthCheckPeriod
	poolConfig.MaxConnIdleTime = maxConnIdleTime
	poolConfig.MaxConnLifetime = maxConnLifetime
	poolConfig.MinConns = minConns

	// Create a new connection pool
	dbPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
