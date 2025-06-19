package controller

import "moment-mail-server/internal/inbox/model"

type InboxUseCase interface {
	CreateInbox() (model.Inbox, error)
}
