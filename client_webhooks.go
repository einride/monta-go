package monta

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
)

const webhooksConfigBasePath = "/v1/webhooks/config"

// WebhookConfig is the request for and response from Get and Update WebhookConfig methods.
type WebhookConfig struct {
	// A HTTPS URL to send the webhook payload to when an event occurs.
	WebhookURL string `json:"webhookUrl"`
	// A cryptoghrapic secret used to sign the webhook payload.
	WebhookSecret string `json:"webhookSecret"`
	// A list of event type to subscribe to. Use ["*"] to subscribe to all.
	EventTypes []*WebhookEventType `json:"eventTypes"`
}

// GetWebhookConfig to get your webhook config.
func (c *clientImpl) GetWebhookConfig(ctx context.Context) (*WebhookConfig, error) {
	query := url.Values{}
	return doGet[WebhookConfig](ctx, c, webhooksConfigBasePath, query)
}

// UpdateWebhookConfig to update your webhook config.
func (c *clientImpl) UpdateWebhookConfig(
	ctx context.Context,
	request *WebhookConfig,
) (*WebhookConfig, error) {
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&request); err != nil {
		return nil, err
	}
	return doPut[WebhookConfig](ctx, c, webhooksConfigBasePath, &requestBody)
}

// DeleteWebhookConfig to delete a webhook config.
func (c *clientImpl) DeleteWebhookConfig(ctx context.Context) error {
	return doDelete(ctx, c, webhooksConfigBasePath)
}
