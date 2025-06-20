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

func (ir *InboxRepository) GetEmailSummaries(inboxId string, limit int, offset int) ([]model.Email, error) {
	const query = `
        SELECT id, sender, subject, received_at
        FROM emails
        WHERE inbox_id = $1
        ORDER BY received_at DESC
		LIMIT $2 OFFSET $3
    `

	rows, err := ir.connection.Query(context.Background(), query, inboxId, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get emails from inbox %s: %w", inboxId, err)
	}
	defer rows.Close()

	summaries := make([]model.Email, 0)

	for rows.Next() {
		var s model.Email
		if err := rows.Scan(
			&s.ID,
			&s.Sender,
			&s.Subject,
			&s.RecievedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan email summary: %w", err)
		}
		summaries = append(summaries, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating email rows: %w", err)
	}

	return summaries, nil
}
