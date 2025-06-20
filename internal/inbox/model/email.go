package model

import (
	"time"

	"github.com/google/uuid"
)

type Email struct {
	ID         uuid.UUID `db:"id"`
	Sender     string    `db:"sender"`
	Subject    string    `db:"subject"`
	RecievedAt time.Time `db:"received_at"`
	Body       string    `db:"body"`
}
