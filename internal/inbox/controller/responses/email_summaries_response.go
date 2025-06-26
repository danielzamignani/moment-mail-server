package responses

type EmailSummariesResponse struct {
	Page           int            `json:"page"`
	Limit          int            `json:"limit"`
	EmailSummaries []EmailSummary `json:"emailSummaries"`
}
