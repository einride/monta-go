package monta

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// StopCharge stop a charge.
func (c *Client) StopCharge(ctx context.Context, chargeID int64) (_ *Charge, err error) {
	method, path := http.MethodPost, fmt.Sprintf("/v1/charges/%d/stop", chargeID)
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
	var response Charge
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
