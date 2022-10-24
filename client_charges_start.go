package monta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
func (c *Client) StartCharge(
	ctx context.Context,
	request *StartChargeRequest,
) (_ *StartChargeResponse, err error) {
	const method, path = http.MethodPost, "/v1/charges" 
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s %s: %w", method, path, err)
		}
	}()
    var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&request); err != nil {
		return nil, err
	}
	requestURL, err := url.Parse(apiHost + path)
	if err != nil {
		return nil, err
	}
	httpRequest, err := http.NewRequestWithContext(ctx, method, requestURL.String(), &requestBody) 
	if err != nil {
		return nil, err
	}
	if err := c.setAuthorization(ctx, httpRequest); err != nil {
		return nil, err
	}
	httpRequest.Header.Set("content-type", "application/json")
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = httpResponse.Body.Close()
	}()
	if httpResponse.StatusCode != http.StatusOK {
		return nil, newStatusError(httpResponse)
	}
	var response StartChargeResponse
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
