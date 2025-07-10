package email

import (
	"context"
	"fmt"

	"github.com/danielzamignani/moment-mail-server/internal/infra/database/postgres"
	"github.com/google/uuid"
)

type EmailRepository struct {
	database *postgres.Database
}

func NewEmailRepository(database *postgres.Database) *EmailRepository {
	return &EmailRepository{
		database: database,
	}
}

func (emailRepository *EmailRepository) GetEmailsSummaries(ctx context.Context, inboxID uuid.UUID) ([]Email, error) {
	query := `
		SELECT id, subject, sender
		FROM emails
		WHERE inbox_id = $1
		ORDER BY received_at DESC
	`
	rows, err := emailRepository.database.Pool.Query(ctx, query, inboxID)
	if err != nil {
		return nil, fmt.Errorf("failed to get emails summaries: %v", err)
	}
	defer rows.Close()

	var emails []Email
	for rows.Next() {
		var email Email
		err := rows.Scan(
			&email.ID,
			&email.Subject,
			&email.Sender,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to get emails summaries: %v", err)
		}

		emails = append(emails, email)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get emails summaries: %v", err)
	}

	return emails, nil
}

func (emailRepository *EmailRepository) GetEmail(ctx context.Context, inboxID uuid.UUID, emailID uuid.UUID) (Email, error) {
	query := `
		SELECT *
		FROM emails
		WHERE inbox_id = $1
		AND id = $2
	`
	var email Email
	err := emailRepository.database.Pool.QueryRow(ctx, query, inboxID, emailID).Scan(
		&email.ID,
		&email.Sender,
		&email.Subject,
		&email.ReceivedAt,
		&email.InboxID,
		&email.Body,
	)

	if err != nil {
		return Email{}, fmt.Errorf("failed to get email in database: %v", err)
	}

	return email, nil
}
