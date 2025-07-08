package service

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/danielzamignani/moment-mail-server/internal/app/dto"
	"github.com/danielzamignani/moment-mail-server/internal/domain/inbox"
	"github.com/google/uuid"
)

type InboxService struct {
	inboxRepository *inbox.InboxRepository
}

func NewInboxService(inboxRepository *inbox.InboxRepository) *InboxService {
	return &InboxService{
		inboxRepository: inboxRepository,
	}
}

func (inboxService *InboxService) CreateInbox(ctx context.Context) (dto.InboxResponse, error) {
	inboxID, err := uuid.NewRandom()
	if err != nil {
		return dto.InboxResponse{}, fmt.Errorf("error generating inbox id: %v", err)
	}

	emailAddess := inboxService.generateEmailAddress()

	now := time.Now().UTC()

	inbox := inbox.Inbox{
		ID:           inboxID,
		EmailAddress: emailAddess,
		CreatedAt:    now,
		ExpiresAt:    now.Add(10 * time.Minute),
	}

	err = inboxService.inboxRepository.CreateInbox(ctx, inbox)
	if err != nil {
		return dto.InboxResponse{}, err
	}

	return dto.InboxResponse{
		ID:           inbox.ID,
		EmailAddress: inbox.EmailAddress,
		CreatedAt:    inbox.CreatedAt,
		ExpiresAt:    inbox.ExpiresAt,
	}, nil
}

func (inboxService *InboxService) DeleteInbox(ctx context.Context, inboxID uuid.UUID) error {
	err := inboxService.inboxRepository.DeleteInbox(ctx, inboxID)
	if err != nil {
		return err
	}

	return nil
}

func (inboxService *InboxService) generateEmailAddress() string {
	username := fmt.Sprintf("momentuser%d", rand.Intn(100000))
	domain := os.Getenv("EMAIL_DOMAIN")

	return fmt.Sprintf("%s@%s", username, domain)
}
