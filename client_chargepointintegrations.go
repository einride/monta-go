package monta

import (
	"context"
	"fmt"
)

// GetChargePointIntegration to retrieve a charge point integration.
func (c *clientImpl) GetChargePointIntegration(
	ctx context.Context,
	chargePointIntegrationID int64,
) (*ChargePointIntegration, error) {
	path := fmt.Sprintf("/v1/charge-point-integrations/%d", chargePointIntegrationID)
	return doGet[ChargePointIntegration](ctx, c, path, nil)
}
