package monta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListChargesRequest is the request input to the [Client.ListCharges] method.
type ListChargesRequest struct {
	// Page number to request (starts with 1).
	Page int
	// Number of items PerPage (between 1 and 100, default 10).
	PerPage int
	// TeamID allows to filter list of charges points by a team ID.
	TeamID *int64
	// ChargePointID allows to filter list of charges points by a charge point ID.
	ChargePointID *int64
}

// ListChargesResponse is the response output from the [Client.ListCharges] method.
type ListChargesResponse struct {
	// Charges in the current page.
	Charges []*Charge `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListCharges to retrieve your charge points.
func (c *Client) ListCharges(
	ctx context.Context,
	request *ListChargesRequest,
) (_ *ListChargesResponse, err error) {
	const method, path = http.MethodGet, "/v1/charges"
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
	query := requestURL.Query()
	if request.Page > 0 {
		query.Set("page", strconv.Itoa(request.Page))
	}
	if request.PerPage > 0 {
		query.Set("perPage", strconv.Itoa(request.PerPage))
	}
	if request.TeamID != nil {
		query.Set("teamId", strconv.Itoa(int(*request.TeamID)))
	}
	if request.ChargePointID != nil {
		query.Set("chargePointId", strconv.Itoa(int(*request.ChargePointID)))
	}
	requestURL.RawQuery = query.Encode()
	httpRequest, err := http.NewRequestWithContext(ctx, method, requestURL.String(), &requestBody)
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
	var response ListChargesResponse
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
