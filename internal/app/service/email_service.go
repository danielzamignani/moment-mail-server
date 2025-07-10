package service

import (
	"context"

	"github.com/danielzamignani/moment-mail-server/internal/domain/email"
	"github.com/google/uuid"
)

type EmailService struct {
	emailRepository *email.EmailRepository
}

func NewEmailService(emailRepository *email.EmailRepository) *EmailService {
	return &EmailService{
		emailRepository: emailRepository,
	}
}

func (emailService *EmailService) GetEmailsSummaries(ctx context.Context, inboxID uuid.UUID) ([]email.Email, error) {
	emails, err := emailService.emailRepository.GetEmailsSummaries(ctx, inboxID)
	if err != nil {
		return nil, err
	}

	return emails, nil
}

func (emailService *EmailService) GetEmail(ctx context.Context, inboxID uuid.UUID, emailID uuid.UUID) (email.Email, error) {
	emailObj, err := emailService.emailRepository.GetEmail(ctx, inboxID, emailID)
	if err != nil {
		return email.Email{}, err
	}

	return emailObj, nil
}
