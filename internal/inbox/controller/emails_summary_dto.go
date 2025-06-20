package controller

import "moment-mail-server/internal/inbox/model"

type EmailsSUmmaryResponse struct {
	Page   int                  `json:"page"`
	Limit  int                  `json:"limit"`
	Emails []model.EmailSummary `json:"emails"`
}
