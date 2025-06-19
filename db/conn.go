package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	encodedPassword := url.QueryEscape(password)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, encodedPassword, host, port, dbname)

	dbpool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		panic(fmt.Sprintf("Failed to ping the database: %v", err))
	}

	log.Println("PostgreSQL database connection established successfully!")

	return dbpool
}
