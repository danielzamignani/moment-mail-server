package responses

import (
	"time"
	"github.com/google/uuid"
)

type EmailSummary struct {
	ID         uuid.UUID `json:"id"`
	Subject    string    `json:"subject"`
	Sender     string    `json:"sender"`
	RecievedAt time.Time `json:"recieved_at"`
}
