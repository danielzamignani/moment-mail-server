package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/danielzamignani/moment-mail-server/internal/config"
	"github.com/danielzamignani/moment-mail-server/internal/infra/database/postgres"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("WARN: .env file not found, using system environment variables")
	}

	config := config.Load()

	_, err := postgres.NewDatabase(context.Background(), config.Database)

	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", ExampleHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Server.Port),
		Handler: mux,
	}

	log.Println("Server is running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error: could not start server: %v", err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test")
}
