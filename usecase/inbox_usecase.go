package usecase

import (
	"fmt"
	"math/rand"
	"moment-mail-server/model"
	"moment-mail-server/repository"
	"time"

	"github.com/google/uuid"
)

type InboxUseCase struct {
	repository *repository.InboxRepository
}

func NewInboxUseCase(repository *repository.InboxRepository) InboxUseCase {
	return InboxUseCase{
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
