package dto

import (
	"time"

	"github.com/google/uuid"
)

type Email struct {
	ID         uuid.UUID `json:"id"`
	Sender     string    `json:"sender"`
	Subject    string    `json:"subject"`
	RecievedAt time.Time `json:"recievedAt"`
	Body       string    `json:"body"`
}
