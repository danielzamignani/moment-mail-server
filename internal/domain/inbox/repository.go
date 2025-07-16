package inbox

import (
	"context"
	"fmt"

	"github.com/danielzamignani/moment-mail-server/internal/infra/database/postgres"
	"github.com/google/uuid"
)

type InboxRepository struct {
	database *postgres.Database
}

func NewInboxRepository(database *postgres.Database) *InboxRepository {
	return &InboxRepository{
		database: database,
	}
}

func (inboxRepository *InboxRepository) CreateInbox(ctx context.Context, inbox Inbox) error {
	query := `INSERT INTO inboxes (id, email_address, created_at, expires_at) 
	VALUES ($1,	$2,	$3, $4)`

	_, err := inboxRepository.database.Pool.Exec(
		ctx,
		query,
		inbox.ID,
		inbox.EmailAddress,
		inbox.CreatedAt,
		inbox.ExpiresAt,
	)

	if err != nil {
		return fmt.Errorf("failed to insert new inbox into the database: %v", err)
	}

	return nil
}

func (inboxRepository *InboxRepository) DeleteInbox(ctx context.Context, inboxID uuid.UUID) error {
	query := `DELETE FROM inboxes WHERE id = $1`

	_, err := inboxRepository.database.Pool.Exec(ctx, query, inboxID)
	if err != nil {
		return fmt.Errorf("failed to delete inbox")
	}

	return nil
}

func (inboxRepository *InboxRepository) GetInboxByID(ctx context.Context, inboxID uuid.UUID) (Inbox, error) {
	query := `
		SELECT *
		FROM inboxes
		WHERE id = $1
	`
	var inbox Inbox
	err := inboxRepository.database.Pool.QueryRow(ctx, query, inboxID).Scan(
		&inbox.ID,
		&inbox.EmailAddress,
		&inbox.ExpiresAt,
		&inbox.CreatedAt,
	)

	if err != nil {
		return Inbox{}, fmt.Errorf("failed to get inbox in database: %v", err)
	}

	return inbox, nil
}
