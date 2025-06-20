package controller

import (
	"moment-mail-server/internal/inbox/dto"
	"moment-mail-server/internal/inbox/model"
)

type InboxUseCase interface {
	CreateInbox() (model.Inbox, error)
	GetEmailsByInboxId(inboxId string, limit int, offset int) ([]dto.EmailSummary, error)
}
