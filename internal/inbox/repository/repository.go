package repository

import (
	"context"
	"fmt"
	"moment-mail-server/internal/inbox/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type InboxRepository struct {
	connection *pgxpool.Pool
}

func NewInboxRepository(connection *pgxpool.Pool) *InboxRepository {
	return &InboxRepository{
		connection: connection,
	}
}

func (ir *InboxRepository) CreateInbox(inbox model.Inbox) error {
	query := `INSERT INTO inboxes (id, email_address, created_at, expires_at) 
	          VALUES ($1, $2, $3, $4)`

	_, err := ir.connection.Exec(context.Background(), query, inbox.ID, inbox.Address, inbox.CreatedAt, inbox.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to insert address into the database: %w", err)
	}

	return nil
}
