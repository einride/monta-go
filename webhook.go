package monta

// WebhookRequest is a webhook request from Monta.
type WebhookRequest struct {
	// List of webhook entries delivered in a batch.
	Entries []WebhookEntry `json:"entries"`

	// Number of pending events left for delivery.
	Pending int64 `json:"pending"`

	// Timestamp of the request (milliseconds from the epoch of 1970-01-01T00:00:00Z).
	Timestamp UnixTimestamp `json:"timestamp"`
}

// WebhookEntry is an entry of the webhook request.
type WebhookEntry struct {
	// Type of the entity.
	EntityType WebhookEntityType `json:"entityType"`

	// ID of the entity.
	EntityID string `json:"entityId"`

	// Type of event, ie. created, deleted, updated.
	EventType WebhookEventType `json:"eventType"`

	// Payload of this entity, e.g. a full Charge object.
	Payload interface{} `json:"payload"`
}
