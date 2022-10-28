package monta

import (
	"context"
	"net/url"
)

// ListSitesRequest is the request input to the [Client.ListSites] method.
type ListSitesRequest struct {
	PageFilters
}

// ListSitesResponse is the response output from the [Client.ListSites] method.
type ListSitesResponse struct {
	// Sites in the current page.
	Sites []*Site `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListSites to retrieve your (charge) sites.
func (c *Client) ListSites(ctx context.Context, request *ListSitesRequest) (*ListSitesResponse, error) {
	path := "/v1/sites"
	query := url.Values{}
	request.PageFilters.Apply(query)
	return doGet[ListSitesResponse](ctx, c, path, query)
}
