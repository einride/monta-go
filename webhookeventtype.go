package monta

// WebhookEventType represents the type of a webhook event.
type WebhookEventType string

// Known [WebhookEventType] values.
const (
	WebhookEventTypeCreated WebhookEventType = "created"
	WebhookEventTypeDeleted WebhookEventType = "deleted"
	WebhookEventTypeUpdated WebhookEventType = "updated"
)
