package responses

import (
	"time"
	"github.com/google/uuid"
)

type InboxResponse struct {
	ID        uuid.UUID `json:"id"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
