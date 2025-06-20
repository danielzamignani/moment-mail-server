package controller

import (
	"moment-mail-server/internal/inbox/dto"
	"moment-mail-server/internal/inbox/model"
)

type InboxService interface {
	CreateInbox() (model.Inbox, error)
	GetEmailSummaries(inboxId string, limit int, offset int) ([]dto.EmailSummary, error)
	GetEmail(emailId string) (dto.Email, error)
}
