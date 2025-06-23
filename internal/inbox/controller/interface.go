package controller

import (
	"moment-mail-server/internal/inbox/dto"
	"moment-mail-server/internal/inbox/model"

	"github.com/google/uuid"
)

type InboxService interface {
	CreateInbox() (model.Inbox, error)
	GetEmailSummaries(inboxId uuid.UUID, limit int, offset int) ([]dto.EmailSummary, error)
	GetEmail(inboxId uuid.UUID, emailId uuid.UUID) (dto.Email, error)
}
