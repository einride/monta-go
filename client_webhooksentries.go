package monta

import (
	"context"
	"net/url"
)

const webhooksEntriesBasePath = "/v1/webhooks/entries"

// ListWebhookEntriesRequest is the request for [Client.ListWebhookEntries] method.
//
// Status filter is not implemented due to undefined integer value mapping by docs at this moment in time.
type ListWebhookEntriesRequest struct {
	PageFilters
}

// ListWebhookEntriesResponse is the response for [Client.ListWebhookEntries] method.
type ListWebhookEntriesResponse struct {
	// List of webhook entries events for your consumer in the past 24 hours
	Events []WebhookEvent `json:"data"`

	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListWebhookEntries to list your webhook entries from the past 24 hours.
func (c *clientImpl) ListWebhookEntries(
	ctx context.Context,
	request *ListWebhookEntriesRequest,
) (*ListWebhookEntriesResponse, error) {
	query := url.Values{}
	request.PageFilters.Apply(query)
	return doGet[ListWebhookEntriesResponse](ctx, c, webhooksEntriesBasePath, query)
}
