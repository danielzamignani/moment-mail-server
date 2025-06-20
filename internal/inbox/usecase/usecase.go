package usecase

import (
	"fmt"
	"math/rand"
	"moment-mail-server/internal/inbox/model"
	"time"

	"github.com/google/uuid"
)

type InboxUseCase struct {
	repository InboxRepository
}

func NewInboxUseCase(repository InboxRepository) *InboxUseCase {
	return &InboxUseCase{
		repository: repository,
	}
}

func (iu *InboxUseCase) CreateInbox() (model.Inbox, error) {
	uuid, _ := uuid.NewUUID()
	emailUsername := fmt.Sprintf("tempuser%d", rand.Intn(100000))
	domainName := "momentmail.com"
	newEmail := fmt.Sprintf("%s@%s", emailUsername, domainName)

	res := model.Inbox{
		ID:        uuid,
		Address:   newEmail,
		CreatedAt: time.Now().UTC(),
		ExpiresAt: time.Now().UTC().Add(10 * time.Minute),
	}

	err := iu.repository.CreateInbox(res)
	if err != nil {
		return model.Inbox{}, err
	}

	return res, nil
}

func (iu *InboxUseCase) GetEmailsByInboxId(inboxId string, limit int, offset int) ([]model.EmailSummary, error) {
	emails, err := iu.repository.GetEmailsByInboxId(inboxId, limit, offset)
	if err != nil {
		return []model.EmailSummary{}, err
	}

	return emails, nil
}
