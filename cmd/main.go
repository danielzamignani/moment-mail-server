package main

import (
	"log"
	"moment-mail-server/controller"
	"moment-mail-server/db"
	"moment-mail-server/repository"
	"moment-mail-server/usecase"
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
	InboxUseCase := usecase.NewInboxUseCase(InboxRepository)
	InboxController := controller.NewInboxController(InboxUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/inbox", InboxController.CreateInbox)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server is running on port", 8080)
	server.ListenAndServe()

}
