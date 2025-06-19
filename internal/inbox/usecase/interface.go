package usecase

import "moment-mail-server/internal/inbox/model"

type InboxRepository interface {
	CreateInbox(inbox model.Inbox) error
	GetEmailsByInboxId(inboxId string) ([]model.EmailSummary, error)
}
