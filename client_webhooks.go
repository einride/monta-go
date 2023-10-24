package monta

import (
	"bytes"
	"context"
	"encoding/json"
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

// UpdateWebhookConfigRequest is the request input to the [Client.UpdateWebhookConfig] method.
type UpdateWebhookConfigRequest struct {
	// A HTTPS URL to send the webhook payload to when an event occurs.
	WebhookURL string `json:"webhookUrl"`
	// A cryptoghrapic secret used to sign the webhook payload.
	WebhookSecret string `json:"webhookSecret"`
	// A list of event type to subscribe to. Use ["*"] to subscribe to all.
	EventTypes []*WebhookEventType `json:"eventTypes"`
}

// UpdateWebhookConfigResponse is the response output from the [Client.UpdateWebhookConfig] method.
type UpdateWebhookConfigResponse struct {
	// A HTTPS URL to send the webhook payload to when an event occurs.
	WebhookURL string `json:"webhookUrl"`
	// A cryptoghrapic secret used to sign the webhook payload.
	WebhookSecret string `json:"webhookSecret"`
	// A list of event type to subscribe to. Use of ["*"] means subscribe to all.
	EventTypes []*string `json:"eventTypes"`
}

// UpdateWebhookConfig to update your webhook config.
func (c *clientImpl) UpdateWebhookConfig(
	ctx context.Context,
	request *UpdateWebhookConfigRequest,
) (*UpdateWebhookConfigResponse, error) {
	path := "/v1/webhooks/config"
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&request); err != nil {
		return nil, err
	}
	return doPut[UpdateWebhookConfigResponse](ctx, c, path, &requestBody)
}
