package responses

import "moment-mail-server/internal/inbox/dto"

type EmailSummariesResponse struct {
	Page           int                `json:"page"`
	Limit          int                `json:"limit"`
	EmailSummaries []dto.EmailSummary `json:"emailSummaries"`
}
