package email

import (
	"time"

	"github.com/google/uuid"
)

type Email struct {
	ID         uuid.UUID
	Sender     string
	Subject    string
	ReceivedAt time.Time
	InboxID    uuid.UUID
	Body       string
}
