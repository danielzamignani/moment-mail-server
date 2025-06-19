package model

import (
	"time"

	"github.com/google/uuid"
)

type EmailSummary struct {
	ID         uuid.UUID `json:"id"`
	Sender     string    `json:"sender"`
	Subject    string    `json:"subject"`
	RecievedAt time.Time `json:"recievedAt"`
}
