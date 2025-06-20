package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type inboxController struct {
	inboxUseCase InboxUseCase
}

func NewInboxController(usecase InboxUseCase) inboxController {
	return inboxController{
		inboxUseCase: usecase,
	}
}

func (i *inboxController) CreateInbox(w http.ResponseWriter, r *http.Request) {
	inbox, err := i.inboxUseCase.CreateInbox()
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

func (i *inboxController) GetEmailsByInboxId(w http.ResponseWriter, r *http.Request) {
	inboxId := r.PathValue("id")
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

	res, err := i.inboxUseCase.GetEmailsByInboxId(inboxId, limit, offset)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})

		return
	}

	response := EmailsSUmmaryResponse{
		Page:   page,
		Limit:  limit,
		Emails: res,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
