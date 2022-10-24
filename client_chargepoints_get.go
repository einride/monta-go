package monta

import (
	"context"
	"fmt"
)

// GetChargePoint to retrieve a single charge point.
func (c *Client) GetChargePoint(ctx context.Context, chargePointID int64) (*ChargePoint, error) {
	path := fmt.Sprintf("/v1/charge-points/%d", chargePointID)
	return doGet[ChargePoint](ctx, c, path, nil)
}
