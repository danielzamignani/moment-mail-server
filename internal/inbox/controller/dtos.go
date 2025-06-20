package controller

import "github.com/google/uuid"

type emailSummary struct {
	ID         uuid.UUID `json:"id"`
	Sender     string    `json:"sender"`
	Subject    string    `json:"subject"`
	RecievedAt uuid.Time `json:"recievedAt"`
}

type EmailSummaries struct {
	EmailSummaries []emailSummary
	Page           int
	Limit          int
}
