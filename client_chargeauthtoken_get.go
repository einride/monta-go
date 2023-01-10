package monta

import (
	"context"
	"fmt"
)

// GetChargeAuthToken to retrieve a single charge auth token.
func (c *Client) GetChargeAuthToken(ctx context.Context, chargeAuthTokenID int64) (*ChargeAuthToken, error) {
	path := fmt.Sprintf("/v1/chargeAuthTokens/%d", chargeAuthTokenID)
	return doGet[ChargeAuthToken](ctx, c, path, nil)
}
