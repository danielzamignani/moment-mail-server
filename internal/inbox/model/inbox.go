package model

import (
	"time"

	"github.com/google/uuid"
)

type Inbox struct {
	ID        uuid.UUID `json:"id"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}
