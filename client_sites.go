package monta

import (
	"context"
	"fmt"
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
func (c *clientImpl) ListSites(ctx context.Context, request *ListSitesRequest) (*ListSitesResponse, error) {
	path := "/v1/sites"
	query := url.Values{}
	request.PageFilters.Apply(query)
	return doGet[ListSitesResponse](ctx, c, path, query)
}

// GetSite to retrieve a single (charge) site.
func (c *clientImpl) GetSite(ctx context.Context, siteID int64) (*Site, error) {
	path := fmt.Sprintf("/v1/sites/%d", siteID)
	return doGet[Site](ctx, c, path, nil)
}
