package monta

import (
	"bytes"
	"context"
	"encoding/json"
)

// StartChargeRequest is the request input to the [Client.StartCharge] method.
type StartChargeRequest struct {
	// PayingTeamID is the ID of the team that will be paying for the charge.
	PayingTeamID int64 `json:"payingTeamId"`
	// ChargePointID is the ID of the charge point used for this charge.
	ChargePointID int64 `json:"chargePointId"`
	// ReserveCharge determines whether the charge point will be reserved or start the charge directly.
	ReserveCharge bool `json:"reserveCharge"`
}

// StartChargeResponse is the response output from the [Client.StartCharge] method.
type StartChargeResponse struct {
	// Charge that started.
	Charge Charge `json:"charge"`
}

// StartCharge starts a charge.
func (c *Client) StartCharge(ctx context.Context, request *StartChargeRequest) (*StartChargeResponse, error) {
	path := "/v1/charges"
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&request); err != nil {
		return nil, err
	}
	return doPost[StartChargeResponse](ctx, c, path, &requestBody)
}
