package service

import "moment-mail-server/internal/inbox/model"

type InboxRepository interface {
	CreateInbox(inbox model.Inbox) error
	GetEmailsByInboxId(inboxId string, limit int, offset int) ([]model.Email, error)
}
