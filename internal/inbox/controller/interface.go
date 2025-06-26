package controller

import (
	"context"
	"moment-mail-server/internal/inbox/controller/responses"

	"github.com/google/uuid"
)

type InboxService interface {
	CreateInbox(ctx context.Context) (responses.InboxResponse, error)
	GetEmailSummaries(ctx context.Context, inboxId uuid.UUID, limit int, offset int) (responses.EmailSummariesResponse, error)
	GetEmail(ctx context.Context, inboxId uuid.UUID, emailId uuid.UUID) (responses.EmailResponse, error)
}
