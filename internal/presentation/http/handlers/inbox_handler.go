package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/danielzamignani/moment-mail-server/internal/app/service"
)

type InboxHandler struct {
	inboxService *service.InboxService
}

func NewInboxHandler(inboxService *service.InboxService) *InboxHandler {
	return &InboxHandler{
		inboxService: inboxService,
	}
}

func (inboxHandler *InboxHandler) CreateInbox(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	inboxResponse, err := inboxHandler.inboxService.CreateInbox(ctx)
	if err != nil {
		inboxHandler.handleError(w, err, http.StatusInternalServerError)
		return
	}

	inboxHandler.writeJSONResponse(w, inboxResponse, http.StatusCreated)
}

func (inboxHandler *InboxHandler) writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (inboxHandler *InboxHandler) handleError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
