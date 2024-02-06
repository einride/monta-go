package monta

import "time"

// WebhookEvent wraps a WebhookEntry and metadata on a webhook call.
type WebhookEvent struct {
	// ID of the webhook event.
	ID int64 `json:"id"`

	// ConsumerID of the intended receiver of the webhook.
	ConsumerID int64 `json:"consumerId"`

	// ID of the operator.
	OperatorID int64 `json:"operatorId"`

	// Entity type related to webhook event.
	WebhookEntityPluralType WebhookEntityPluralType `json:"eventType"`

	// Payload of the related webhook call.
	WebhookEntry WebhookEntry `json:"payload"`

	// Status of the webhook call.
	Status WebhookEventStatus `json:"status"`

	// Error related to the webhook call.
	Error string `json:"error"`

	// When the webhook event was created.
	CreatedAt time.Time `json:"createdAt"`

	// When the webhook event was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
}

// Status of the webhook call.
type WebhookEventStatus string

// Known [WebhookEventStatus] values.
const (
	WebhookStatusPending   WebhookEventStatus = "pending"
	WebhookStatusCompleted WebhookEventStatus = "completed"
	WebhookStatusFailure   WebhookEventStatus = "failure"
)
