package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/danielzamignani/moment-mail-server/internal/app/service"
	"github.com/danielzamignani/moment-mail-server/internal/config"
	"github.com/danielzamignani/moment-mail-server/internal/domain/inbox"
	"github.com/danielzamignani/moment-mail-server/internal/infra/database/postgres"
	"github.com/danielzamignani/moment-mail-server/internal/presentation/http/handlers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("WARN: .env file not found, using system environment variables")
	}

	config := config.Load()

	database, err := postgres.NewDatabase(context.Background(), config.Database)

	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	inboxReposiroty := inbox.NewInboxRepository(database)
	inboxService := service.NewInboxService(inboxReposiroty)
	inboxHandler := handlers.NewInboxHandler(inboxService)

	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /v1/inbox/{inboxID}", inboxHandler.DeleteInbox)
	mux.HandleFunc("POST /v1/inbox", inboxHandler.CreateInbox)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Server.Port),
		Handler: mux,
	}

	log.Println("Server is running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error: could not start server: %v", err)
	}
}
