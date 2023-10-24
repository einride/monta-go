package monta

import (
	"context"
	"net/url"
)

// GetWebhookConfigResponse is the response output from the [Client.GetWebhookConfig] method.
type GetWebhookConfigResponse struct {
	// A HTTPS URL to send the webhook payload to when an event occurs.
	WebhookURL string `json:"webhookUrl"`
	// A cryptoghrapic secret used to sign the webhook payload.
	WebhookSecret string `json:"webhookSecret"`
	// A list of event types to subscribe to. Use of ["*"] means subscribe to all.
	EventTypes []*WebhookEventType `json:"eventTypes"`
}

// GetWebhookConfig to get your webhook config.
func (c *clientImpl) GetWebhookConfig(ctx context.Context) (*GetWebhookConfigResponse, error) {
	path := "/v1/webhooks/config"
	query := url.Values{}
	return doGet[GetWebhookConfigResponse](ctx, c, path, query)
}
