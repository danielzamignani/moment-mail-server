package inbox

import (
	"context"
	"fmt"

	"github.com/danielzamignani/moment-mail-server/internal/infra/database/postgres"
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
