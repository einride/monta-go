package monta

import (
	"context"
	"fmt"
	"net/url"
)

// GetChargePoint to retrieve a single charge point
func (c *Client) GetChargePoint(ctx context.Context, chargePointID int64) (_ *ChargePoint, err error) {
	path := fmt.Sprintf("/v1/charge-points/%d", chargePointID)
	return doGet[ChargePoint](ctx, c, path, url.Values{})
}
