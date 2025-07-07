package postgres

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/danielzamignani/moment-mail-server/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase(ctx context.Context, config config.DatabaseConfig) (*Database, error) {
	connString := buildConnectionString(config)

	log.Println("Connecting to database...")

	pool, err := pgxpool.New(ctx, connString)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping to database: %v", err)
	}

	log.Println("connected to database successfully")

	return &Database{Pool: pool}, nil
}

func buildConnectionString(config config.DatabaseConfig) string {
	encodedPassword := url.QueryEscape(config.Password)
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		encodedPassword,
		config.Host,
		config.Port,
		config.DBName,
	)

	return connString
}
