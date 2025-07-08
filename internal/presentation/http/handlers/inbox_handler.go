package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielzamignani/moment-mail-server/internal/app/service"
	"github.com/google/uuid"
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
	ctx := r.Context()

	inboxResponse, err := inboxHandler.inboxService.CreateInbox(ctx)
	if err != nil {
		inboxHandler.handleError(w, err, http.StatusInternalServerError)
		return
	}

	inboxHandler.writeJSONResponse(w, inboxResponse, http.StatusCreated)
}

func (inboxHandler *InboxHandler) DeleteInbox(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	inboxIDStr := r.PathValue("inboxID")
	inboxID, err := uuid.Parse(inboxIDStr)
	if err != nil {
		inboxHandler.handleError(w, fmt.Errorf("failed to parse id param: %v", err), http.StatusBadRequest)
		return
	}

	if err := inboxHandler.inboxService.DeleteInbox(ctx, inboxID); err != nil {
		inboxHandler.handleError(w, fmt.Errorf("failed to delete inbox: %v", err), http.StatusInternalServerError)
		return
	}

	inboxHandler.writeJSONResponse(w, nil, http.StatusNoContent)
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
