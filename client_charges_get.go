package monta

import (
	"context"
	"fmt"
)

// GetCharge to retrieve a single charge
func (c *Client) GetCharge(ctx context.Context, chargeID int64) (_ *Charge, err error) {
	return getEntity[Charge](ctx, c, fmt.Sprintf("/v1/charges/%d", chargeID))
}
