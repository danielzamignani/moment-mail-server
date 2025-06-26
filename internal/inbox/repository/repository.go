package repository

import (
	"context"
	"fmt"
	"moment-mail-server/internal/inbox/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
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

func (ir *InboxRepository) CreateInbox(ctx context.Context, inbox model.Inbox) error {
	query := `INSERT INTO inboxes (id, email_address, created_at, expires_at) 
	          VALUES ($1, $2, $3, $4)`

	_, err := ir.connection.Exec(ctx, query, inbox.ID, inbox.Address, inbox.CreatedAt, inbox.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to insert address into the database: %w", err)
	}

	return nil
}

func (ir *InboxRepository) GetEmailSummaries(ctx context.Context, inboxId uuid.UUID, limit int, offset int) ([]model.Email, error) {
	const query = `
        SELECT id, sender, subject, received_at
        FROM emails
        WHERE inbox_id = $1
        ORDER BY received_at DESC
		LIMIT $2 OFFSET $3
    `

	rows, err := ir.connection.Query(ctx, query, inboxId, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get emails from inbox %s: %w", inboxId, err)
	}
	defer rows.Close()

	var summaries []model.Email

	for rows.Next() {
		var s model.Email
		if err := rows.Scan(
			&s.ID,
			&s.Sender,
			&s.Subject,
			&s.ReceivedAt,
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

func (ir *InboxRepository) GetEmail(ctx context.Context, inboxId uuid.UUID, emailId uuid.UUID) (model.Email, error) {
	const query = `
        SELECT id, received_at, sender, subject, body, inbox_id
        FROM emails
        WHERE id = $1 AND inbox_id = $2
    `

	row := ir.connection.QueryRow(ctx, query, emailId, inboxId)

	var email model.Email
	if err := row.Scan(
		&email.ID,
		&email.ReceivedAt,
		&email.Sender,
		&email.Subject,
		&email.Body,
		&email.InboxID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return model.Email{}, fmt.Errorf("email not found")
		}
		return model.Email{}, fmt.Errorf("repository: failed to scan email: %w", err)
	}

	return email, nil
}
