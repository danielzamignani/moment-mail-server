package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/danielzamignani/moment-mail-server/internal/app/service"
	"github.com/google/uuid"
)

type EmailHandler struct {
	emailService *service.EmailService
}

func NewEmailHandler(emailService *service.EmailService) *EmailHandler {
	return &EmailHandler{
		emailService: emailService,
	}
}

func (emailHandler *EmailHandler) GetEmailsSummaries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	inboxIDStr := r.PathValue("inboxID")
	inboxID, err := uuid.Parse(inboxIDStr)

	emailsResponse, err := emailHandler.emailService.GetEmailsSummaries(ctx, inboxID)
	if err != nil {
		emailHandler.handleError(w, err, http.StatusInternalServerError)
		return
	}

	emailHandler.writeJSONResponse(w, emailsResponse, http.StatusCreated)
}

func (emailHandler *EmailHandler) GetEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	inboxIDStr := r.PathValue("inboxID")
	inboxID, err := uuid.Parse(inboxIDStr)

	emailIDStr := r.PathValue("emailID")
	emailID, err := uuid.Parse(emailIDStr)

	emailResponse, err := emailHandler.emailService.GetEmail(ctx, inboxID, emailID)
	if err != nil {
		emailHandler.handleError(w, err, http.StatusInternalServerError)
		return
	}

	emailHandler.writeJSONResponse(w, emailResponse, http.StatusCreated)
}

func (emailHandler *EmailHandler) writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (emailHandler *EmailHandler) handleError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
