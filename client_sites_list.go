package monta

import (
	"context"
	"net/url"
	"strconv"
)

// ListSitesRequest is the request input to the [Client.ListSites] method.
type ListSitesRequest struct {
	// Page number to request (starts with 1).
	Page int
	// Number of items PerPage (between 1 and 100, default 10).
	PerPage int
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
	if request.Page > 0 {
		query.Set("page", strconv.Itoa(request.Page))
	}
	if request.PerPage > 0 {
		query.Set("perPage", strconv.Itoa(request.PerPage))
	}
	return doGet[ListSitesResponse](ctx, c, path, query)
}
