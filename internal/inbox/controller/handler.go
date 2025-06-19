package controller

import (
	"encoding/json"
	"moment-mail-server/internal/inbox/usecase"
	"net/http"
)

type inboxController struct {
	inboxUseCase usecase.InboxUseCase
}

func NewInboxController(usecase usecase.InboxUseCase) inboxController {
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

	res, err := i.inboxUseCase.GetEmailsByInboxId(inboxId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"emails": res,
	})
}
