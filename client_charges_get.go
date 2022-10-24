package monta

import (
	"context"
	"fmt"
)

// GetCharge to retrieve a single charge.
func (c *Client) GetCharge(ctx context.Context, chargeID int64) (*Charge, error) {
	path := fmt.Sprintf("/v1/charges/%d", chargeID)
	return doGet[Charge](ctx, c, path, nil)
}
