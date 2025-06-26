package main

import (
	"log"
	"moment-mail-server/db"
	"moment-mail-server/internal/broker"
	"moment-mail-server/internal/inbox/controller"
	"moment-mail-server/internal/inbox/repository"
	"moment-mail-server/internal/inbox/service"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("WARN: .env file not found, using system environment variables.")
	}

	dbConnection := db.ConnectDB()
	inboxRepository := repository.NewInboxRepository(dbConnection)
	inboxService := service.NewInboxService(inboxRepository)
	broker := broker.NewEventBroker()
	inboxController := controller.NewInboxController(inboxService, broker)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/inbox", inboxController.CreateInbox)
	mux.HandleFunc("GET /api/inbox/{inboxId}/emails", inboxController.GetEmailSummaries)
	mux.HandleFunc("GET /api/inbox/{inboxId}/emails/{emailId}", inboxController.GetEmail)
	mux.HandleFunc("GET /api/events/{inboxId}", inboxController.EventHandler)
	mux.HandleFunc("POST /test/publish", inboxController.TestPublishHandler)
	mux.HandleFunc("DELETE /api/inbox/{inboxId}", inboxController.DeleteInbox)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server is running on port", 8080)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("FATAL: could not start server: %v", err)
	}
}
