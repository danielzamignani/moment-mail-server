package service

import (
	"context"
	"fmt"
	"math/rand"
	"moment-mail-server/internal/inbox/controller/responses"
	"moment-mail-server/internal/inbox/model"
	"time"

	"github.com/google/uuid"
)

type InboxService struct {
	repository InboxRepository
}

func NewInboxService(repository InboxRepository) *InboxService {
	return &InboxService{
		repository: repository,
	}
}

func (iu *InboxService) CreateInbox(ctx context.Context) (responses.InboxResponse, error) {
	inboxID, err := uuid.NewUUID()
	if err != nil {
		return responses.InboxResponse{}, fmt.Errorf("service: failed to generate UUID: %w", err)
	}

	emailUsername := fmt.Sprintf("tempuser%d", rand.Intn(100000))
	domainName := "momentmail.com"
	newEmail := fmt.Sprintf("%s@%s", emailUsername, domainName)

	res := model.Inbox{
		ID:        inboxID,
		Address:   newEmail,
		CreatedAt: time.Now().UTC(),
		ExpiresAt: time.Now().UTC().Add(10 * time.Minute),
	}

	err = iu.repository.CreateInbox(ctx, res)
	if err != nil {
		return responses.InboxResponse{}, fmt.Errorf("service: failed to create inbox: %w", err)
	}

	return responses.InboxResponse{
		ID:        res.ID,
		Address:   res.Address,
		CreatedAt: res.CreatedAt,
		ExpiresAt: res.ExpiresAt,
	}, nil
}

func (iu *InboxService) GetEmailSummaries(ctx context.Context, inboxId uuid.UUID, limit int, offset int) (responses.EmailSummariesResponse, error) {
	emails, err := iu.repository.GetEmailSummaries(ctx, inboxId, limit, offset)
	if err != nil {
		return responses.EmailSummariesResponse{}, fmt.Errorf("service: failed to get email summaries: %w", err)
	}

	summaries := make([]responses.EmailSummary, len(emails))
	for i, email := range emails {
		summaries[i] = responses.EmailSummary{
			ID:         email.ID,
			Subject:    email.Subject,
			Sender:     email.Sender,
			RecievedAt: email.ReceivedAt,
		}
	}

	return responses.EmailSummariesResponse{
		EmailSummaries: summaries,
	}, nil
}

func (iu *InboxService) GetEmail(ctx context.Context, inboxId uuid.UUID, emailId uuid.UUID) (responses.EmailResponse, error) {
	emailModel, err := iu.repository.GetEmail(ctx, inboxId, emailId)
	if err != nil {
		return responses.EmailResponse{}, err
	}

	return responses.EmailResponse{
		ID:         emailModel.ID,
		Subject:    emailModel.Subject,
		Sender:     emailModel.Sender,
		RecievedAt: emailModel.ReceivedAt,
		Body:       emailModel.Body,
		InboxID:    emailModel.InboxID,
	}, nil
}
