package monta

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// GetChargePoint to retrieve a single charge point
func (c *Client) GetChargePoint(ctx context.Context, chargePointID int64) (_ *ChargePoint, err error) {
	method, path := http.MethodGet, fmt.Sprintf("/v1/charge-points/%d", chargePointID)
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s %s: %w", method, path, err)
		}
	}()
	requestURL, err := url.Parse(apiHost + path)
	if err != nil {
		return nil, err
	}
	httpRequest, err := http.NewRequestWithContext(ctx, method, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	if err := c.setAuthorization(ctx, httpRequest); err != nil {
		return nil, err
	}
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
	var response ChargePoint
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
