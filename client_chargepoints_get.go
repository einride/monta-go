package monta

import (
	"context"
	"fmt"
)

// GetChargePoint to retrieve a single charge point
func (c *Client) GetChargePoint(ctx context.Context, chargePointID int64) (_ *ChargePoint, err error) {
	return getEntity[ChargePoint](ctx, c, fmt.Sprintf("/v1/charge-points/%d", chargePointID))
}
