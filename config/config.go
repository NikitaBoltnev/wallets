package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	Port       string
}

func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func LoadConfigAndPool() (*Config, *pgxpool.Pool, error) {
	cfg := &Config{

		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("SERVER_PORT"),
	}

	if cfg.DBUsername == "" || cfg.DBPassword == "" || cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBName == "" {
		return nil, nil, fmt.Errorf("one or more required environment variables are not set in config.env")
	}

	ctx := context.Background()

	bd := cfg.GetDBConnectionString()

	pool, err := pgxpool.New(ctx, bd)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create database pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return cfg, pool, nil
}
