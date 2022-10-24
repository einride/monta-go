package monta

import (
	"context"
	"fmt"
	"net/url"
)

// RestartCharge restart or start a reserved charge.
func (c *Client) RestartCharge(ctx context.Context, chargeID int64) (*Charge, error) {
	path := fmt.Sprintf("/v1/charges/%d/restart", chargeID)
	return doGet[Charge](ctx, c, path, url.Values{})
}
