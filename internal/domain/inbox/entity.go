package inbox

import (
	"time"

	"github.com/google/uuid"
)

type Inbox struct {
	ID           uuid.UUID
	EmailAddress string
	CreatedAt    time.Time
	ExpiresAt    time.Time
}
