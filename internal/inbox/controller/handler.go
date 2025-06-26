package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"moment-mail-server/internal/broker"
	"moment-mail-server/internal/inbox/controller/responses"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type inboxController struct {
	inboxService InboxService
	broker       *broker.EventBroker
}

func NewInboxController(service InboxService, broker *broker.EventBroker) inboxController {
	return inboxController{
		inboxService: service,
		broker:       broker,
	}
}

func (i *inboxController) CreateInbox(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	inbox, err := i.inboxService.CreateInbox(ctx)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	response := responses.InboxResponse{
		ID:        inbox.ID,
		Address:   inbox.Address,
		CreatedAt: inbox.CreatedAt,
		ExpiresAt: inbox.ExpiresAt,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (i *inboxController) GetEmailSummaries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	inboxIdStr := r.PathValue("inboxId")
	inboxId, err := uuid.Parse(inboxIdStr)
	if err != nil {
		http.Error(w, "Invalid inbox ID", http.StatusBadRequest)
		return
	}
	query := r.URL.Query()
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || limit < 1 || limit > 100 {
		limit = 20
	}
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	response, err := i.inboxService.GetEmailSummaries(ctx, inboxId, limit, offset)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}
	response.Page = page
	response.Limit = limit
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (i *inboxController) GetEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	inboxIdStr := r.PathValue("inboxId")
	emailIdStr := r.PathValue("emailId")

	inboxId, err := uuid.Parse(inboxIdStr)
	if err != nil {
		http.Error(w, "Invalid inbox ID", http.StatusBadRequest)
		return
	}

	emailId, err := uuid.Parse(emailIdStr)
	if err != nil {
		http.Error(w, "Invalid email ID", http.StatusBadRequest)
		return
	}

	response, err := i.inboxService.GetEmail(ctx, inboxId, emailId)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (i *inboxController) EventHandler(w http.ResponseWriter, r *http.Request) {
	inboxIdStr := r.PathValue("inboxId")

	inboxId, err := uuid.Parse(inboxIdStr)
	if err != nil {
		http.Error(w, "Invalid inbox ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	i.broker.Subscribe(inboxId)

	messageChan := i.broker.Subscribe(inboxId)

	defer i.broker.Unsubscribe(inboxId)

	for {
		select {
		case event, ok := <-messageChan:
			if !ok {
				return
			}

			fmt.Fprintf(w, "event: %s\n", event.Type)
			fmt.Fprintf(w, "data: \n\n")

			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}

		case <-r.Context().Done():
			log.Printf("INFO: Client %s desconected.", inboxId)
			return
		}
	}
}

func (i *inboxController) TestPublishHandler(w http.ResponseWriter, r *http.Request) {
	inboxId, _ := uuid.Parse(r.PathValue("inboxId"))

	event := broker.Event{
		Type: "new_email",
	}

	i.broker.Publish(inboxId, event)

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Event publish for client"))
}
