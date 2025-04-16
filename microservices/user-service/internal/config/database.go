package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	PostgresURL string
	SupabaseURL string
	SupabaseKey string
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	return &DatabaseConfig{
		PostgresURL: os.Getenv("POSTGRES_URL"),
		SupabaseURL: os.Getenv("SUPABASE_URL"),
		SupabaseKey: os.Getenv("SUPABASE_KEY"),
	}, nil
}

// NewPostgresConnection crée une nouvelle connexion PostgreSQL
func NewPostgresConnection(ctx context.Context, config *DatabaseConfig) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, config.PostgresURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	// Vérifier la connexion
	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return conn, nil
}
