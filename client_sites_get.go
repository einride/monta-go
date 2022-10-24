package monta

import (
	"context"
	"fmt"
)

// GetSite to retrieve a single (charge) site
func (c *Client) GetSite(ctx context.Context, siteID int64) (_ *Site, err error) {
	return getEntity[Site](ctx, c, fmt.Sprintf("/v1/sites/%d", siteID))
}
