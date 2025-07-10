package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/danielzamignani/moment-mail-server/internal/app/service"
	"github.com/danielzamignani/moment-mail-server/internal/config"
	"github.com/danielzamignani/moment-mail-server/internal/domain/email"
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

	mux := http.NewServeMux()

	inboxReposiroty := inbox.NewInboxRepository(database)
	inboxService := service.NewInboxService(inboxReposiroty)
	inboxHandler := handlers.NewInboxHandler(inboxService)

	emailRepository := email.NewEmailRepository(database)
	emailService := service.NewEmailService(emailRepository)
	emailHandler := handlers.NewEmailHandler(emailService)

	mux.HandleFunc("POST /v1/inbox", inboxHandler.CreateInbox)
	mux.HandleFunc("DELETE /v1/inbox/{inboxID}", inboxHandler.DeleteInbox)
	mux.HandleFunc("GET /v1/inbox/{inboxID}/emails", emailHandler.GetEmailsSummaries)
	mux.HandleFunc("GET /v1/inbox/{inboxID}/emails/{emailID}", emailHandler.GetEmail)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Server.Port),
		Handler: mux,
	}

	log.Println("Server is running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error: could not start server: %v", err)
	}
}
