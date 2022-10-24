package monta

import (
	"context"
	"fmt"
)

// StopCharge stop a charge.
func (c *Client) StopCharge(ctx context.Context, chargeID int64) (_ *Charge, err error) {
	path := fmt.Sprintf("/v1/charges/%d/stop", chargeID)
	return doPost[Charge](ctx, c, path, nil)
}
