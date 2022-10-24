package monta

import (
	"context"
	"fmt"
	"net/url"
)

// GetSite to retrieve a single (charge) site
func (c *Client) GetSite(ctx context.Context, siteID int64) (*Site, error) {
	path := fmt.Sprintf("/v1/sites/%d", siteID)
	return doGet[Site](ctx, c, path, url.Values{})
}
