package models

type PaymentRequest struct {
	PaymentMethod string `json:"payment_method"` // ALL DATA TO PROCESS PAYMENT
	WebhookURL    string `json:"webhook_url"`
	UniqueID      string `json:"unique_id"`
}

type PaymentResponse struct {
	Status   string `json:"status"`
	UniqueID string `json:"unique_id"`
}
