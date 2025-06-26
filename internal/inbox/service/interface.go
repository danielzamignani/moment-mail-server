package service

import (
	"context"
	"moment-mail-server/internal/inbox/model"

	"github.com/google/uuid"
)

type InboxRepository interface {
	CreateInbox(ctx context.Context, inbox model.Inbox) error
	GetEmailSummaries(ctx context.Context, inboxId uuid.UUID, limit int, offset int) ([]model.Email, error)
	GetEmail(ctx context.Context, inboxId uuid.UUID, emailId uuid.UUID) (model.Email, error)
}
