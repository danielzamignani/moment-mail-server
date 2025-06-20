package service

import (
	"fmt"
	"math/rand"
	"moment-mail-server/internal/inbox/dto"
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

func (iu *InboxService) CreateInbox() (model.Inbox, error) {
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

func (iu *InboxService) GetEmailSummaries(inboxId string, limit int, offset int) ([]dto.EmailSummary, error) {
	emails, err := iu.repository.GetEmailSummaries(inboxId, limit, offset)
	if err != nil {
		return []dto.EmailSummary{}, err
	}

	summaries := make([]dto.EmailSummary, len(emails))
	for i, email := range emails {
		summaries[i] = dto.EmailSummary{
			ID:         email.ID,
			Subject:    email.Subject,
			Sender:     email.Sender,
			RecievedAt: email.RecievedAt,
		}
	}

	return summaries, nil
}

func (iu *InboxService) GetEmail(emailId string) (dto.Email, error) {
	emailModel, err := iu.repository.GetEmail(emailId)
	if err != nil {
		return dto.Email{}, err
	}

	email := dto.Email{
		ID:         emailModel.ID,
		Subject:    emailModel.Subject,
		Sender:     emailModel.Sender,
		RecievedAt: emailModel.RecievedAt,
		Body:       emailModel.Body,
	}

	return email, nil
}
