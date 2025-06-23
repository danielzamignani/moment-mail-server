package controller

import (
	"encoding/json"
	"moment-mail-server/internal/inbox/controller/responses"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type inboxController struct {
	inboxService InboxService
}

func NewInboxController(service InboxService) inboxController {
	return inboxController{
		inboxService: service,
	}
}

func (i *inboxController) CreateInbox(w http.ResponseWriter, r *http.Request) {
	inbox, err := i.inboxService.CreateInbox()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(inbox)
}

func (i *inboxController) GetEmailSummaries(w http.ResponseWriter, r *http.Request) {
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

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	res, err := i.inboxService.GetEmailSummaries(inboxId, limit, offset)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})

		return
	}

	response := responses.EmailSummariesResponse{
		Page:           page,
		Limit:          limit,
		EmailSummaries: res,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (i *inboxController) GetEmail(w http.ResponseWriter, r *http.Request) {
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

	res, err := i.inboxService.GetEmail(inboxId, emailId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})

		return
	}

	response := res

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
