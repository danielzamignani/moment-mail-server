package main

import (
	"log"
	"moment-mail-server/db"
	"moment-mail-server/internal/inbox/controller"
	"moment-mail-server/internal/inbox/repository"
	"moment-mail-server/internal/inbox/service"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	dbConnection := db.ConnectDB()
	InboxRepository := repository.NewInboxRepository(dbConnection)
	InboxService := service.NewInboxService(InboxRepository)
	InboxController := controller.NewInboxController(InboxService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/inbox", InboxController.CreateInbox)
	mux.HandleFunc("GET /api/inbox/{id}/emails", InboxController.GetEmailSummaries)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server is running on port", 8080)
	server.ListenAndServe()

}
