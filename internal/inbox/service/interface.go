package service

import (
	"moment-mail-server/internal/inbox/model"

	"github.com/google/uuid"
)

type InboxRepository interface {
	CreateInbox(inbox model.Inbox) error
	GetEmailSummaries(inboxId uuid.UUID, limit int, offset int) ([]model.Email, error)
	GetEmail(inboxId uuid.UUID, emailId uuid.UUID) (model.Email, error)
}
