package monta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// CreateChargeAuthTokenRequest is the request input to the [Client.CreateChargeAuthToken] method.
type CreateChargeAuthTokenRequest struct {
	// Id of the team the charge auth token belongs.
	TeamID int64 `json:"teamId"`
	// Id of the user the charge auth token should be associated to.
	UserID *int64 `json:"userId"`
	// Identifier of the the charge auth token. Note: without prefix e.g VID:
	Identifier string `json:"identifier"`
	// Type of the charge auth token.
	Type ChargeAuthTokenType `json:"type"`
	// Name of the charge auth token.
	Name *string `json:"name"`
	// If the charge auth token should be active in the Monta network.
	MontaNetwork bool `json:"montaNetwork"`
	// If the charge auth token should be active in the Roaming network.
	RoamingNetwork bool `json:"roamingNetwork"`
	// Until when the charge auth token should be active, when null it will be active forever.
	ActiveUntil *time.Time `json:"activeUntil"`
}

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

// CreateChargeAuthToken to create a new charge auth token.
func (c *clientImpl) CreateChargeAuthToken(
	ctx context.Context,
	request CreateChargeAuthTokenRequest,
) (*ChargeAuthToken, error) {
	path := "/v1/charge-auth-tokens"
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&request); err != nil {
		return nil, err
	}
	return doPost[ChargeAuthToken](ctx, c, path, &requestBody)
}

// ListChargeAuthTokens to retrieve your charge auth tokens.
func (c *clientImpl) ListChargeAuthTokens(
	ctx context.Context,
	request *ListChargeAuthTokensRequest,
) (*ListChargeAuthTokensResponse, error) {
	path := "/v1/charge-auth-tokens"
	query := url.Values{}
	request.PageFilters.Apply(query)
	if request.TeamID != nil {
		query.Set("teamId", strconv.Itoa(int(*request.TeamID)))
	}
	return doGet[ListChargeAuthTokensResponse](ctx, c, path, query)
}

// GetChargeAuthToken to retrieve a single charge auth token.
func (c *clientImpl) GetChargeAuthToken(ctx context.Context, chargeAuthTokenID int64) (*ChargeAuthToken, error) {
	path := fmt.Sprintf("/v1/charge-auth-tokens/%d", chargeAuthTokenID)
	return doGet[ChargeAuthToken](ctx, c, path, nil)
}
