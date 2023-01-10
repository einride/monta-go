package monta

import (
	"context"
	"net/url"
	"strconv"
)

// ListChargeAuthTokensRequest is the request input to the [Client.ListChargeAuthTokens] method.
type ListChargeAuthTokensRequest struct {
	PageFilters
	// Filter to retrieve charges auth tokens with specified teamId.
	TeamID *int64
}

// ListChargeAuthTokensResponse is the response output from the [Client.ListChargeAuthTokens] method.
type ListChargeAuthTokensResponse struct {
	// Charges in the current page.
	ChargeAuthTokens []*ChargeAuthToken `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListChargeAuthTokens to retrieve your charge auth tokens.
func (c *Client) ListChargeAuthTokens(
	ctx context.Context,
	request *ListChargeAuthTokensRequest,
) (*ListChargeAuthTokensResponse, error) {
	path := "/v1/chargeAuthTokens"
	query := url.Values{}
	request.PageFilters.Apply(query)
	if request.TeamID != nil {
		query.Set("teamId", strconv.Itoa(int(*request.TeamID)))
	}
	return doGet[ListChargeAuthTokensResponse](ctx, c, path, query)
}
