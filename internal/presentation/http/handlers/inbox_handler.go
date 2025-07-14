package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/danielzamignani/moment-mail-server/internal/app/service"
	"github.com/danielzamignani/moment-mail-server/internal/broker"
	"github.com/google/uuid"
)

type InboxHandler struct {
	inboxService *service.InboxService
	broker       *broker.EventBroker
}

func NewInboxHandler(inboxService *service.InboxService, broker *broker.EventBroker) *InboxHandler {
	return &InboxHandler{
		inboxService: inboxService,
		broker:       broker,
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

func (inboxHandler *InboxHandler) EventHandler(w http.ResponseWriter, r *http.Request) {
	inboxIDStr := r.PathValue("inboxID")
	inboxID, err := uuid.Parse(inboxIDStr)

	if err != nil {
		inboxHandler.handleError(w, fmt.Errorf("failed to parse id param: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	messageChan := inboxHandler.broker.Subscribe(inboxID)
	defer inboxHandler.broker.Unsubscribe(inboxID)

	for {
		select {
		case event, ok := <-messageChan:
			if !ok {
				return
			}

			eventBytes, err := json.Marshal(event.Data)
			if err != nil {
				log.Printf("Error to serialize event data %s: %v", inboxID, err)
				continue
			}

			fmt.Fprintf(w, "event: %s\n", event.Type)
			fmt.Fprintf(w, "data: %s\n\n", eventBytes)

			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		case <-r.Context().Done():
			log.Printf("Client %s disconected", inboxID)
			return
		}
	}
}

func (inboxHandler *InboxHandler) TestPublishHandler(w http.ResponseWriter, r *http.Request) {
	inboxIDStr := r.PathValue("inboxID")
	inboxID, err := uuid.Parse(inboxIDStr)

	if err != nil {
		inboxHandler.handleError(w, fmt.Errorf("failed to parse id param: %v", err), http.StatusBadRequest)
		return
	}

	event := broker.Event{
		Type: "new_email",
		Data: map[string]string{
			"message": "You have a new email!",
			"emailId": uuid.New().String(),
		},
	}

	inboxHandler.broker.Publish(inboxID, event)
	log.Printf("Event test publish to inboxID: %s", inboxID)

	inboxHandler.writeJSONResponse(w, map[string]string{"message": "Event test publish to " + inboxIDStr}, http.StatusOK)
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
