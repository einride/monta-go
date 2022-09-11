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

// ListChargePointsRequest is the request input to the [Client.ListChargePoints] method.
type ListChargePointsRequest struct {
	// Page number to request (starts with 1).
	Page int
	// Number of items PerPage (between 1 and 100, default 10).
	PerPage int
	// SiteID allows to filter list of charge points by a site id.
	SiteID *int64
}

// ListChargePointsResponse is the response output from the [Client.ListChargePoints] method.
type ListChargePointsResponse struct {
	// ChargePoints in the current page.
	ChargePoints []*ChargePoint `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListChargePoints to retrieve your charge points.
func (c *Client) ListChargePoints(
	ctx context.Context,
	request *ListChargePointsRequest,
) (_ *ListChargePointsResponse, err error) {
	const method, path = http.MethodGet, "/v1/charge-points"
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
	if request.SiteID != nil {
		query.Set("siteId", strconv.Itoa(int(*request.SiteID)))
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
	var response ListChargePointsResponse
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
