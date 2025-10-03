package monta

import (
	"context"
	"fmt"
)

// GetChargePointModel to retrieve a single charge point model.
func (c *clientImpl) GetChargePointModel(ctx context.Context, chargePointModelID int64) (*ChargePointModel, error) {
	path := fmt.Sprintf("/v1/charge-point-models/%d", chargePointModelID)
	return doGet[ChargePointModel](ctx, c, path, nil)
}
