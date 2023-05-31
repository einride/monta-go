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

//nolint:gosec
const chargeAuthTokenBasePath = "/v1/charge-auth-tokens"

// CreateChargeAuthTokenRequest is the request input to the [Client.CreateChargeAuthToken] method.
type CreateChargeAuthTokenRequest struct {
	// Id of the team the charge auth token belongs to.
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
	// Until when the charge auth token should be active, when nil it will be active forever.
	ActiveUntil *time.Time `json:"activeUntil"`
}

// PatchChargeAuthTokenRequest is the request input to the [Client.PatchChargeAuthToken] method.
type PatchChargeAuthTokenRequest struct {
	// External Id of this entity, managed by you.
	PartnerExternalID *string `json:"partnerExternalId,omitempty"`
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
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&request); err != nil {
		return nil, err
	}
	return doPost[ChargeAuthToken](ctx, c, chargeAuthTokenBasePath, &requestBody)
}

// PatchChargeAuthToken to patch a charge auth token.
func (c *clientImpl) PatchChargeAuthToken(
	ctx context.Context,
	chargeAuthTokenID int64,
	request PatchChargeAuthTokenRequest,
) (*ChargeAuthToken, error) {
	path := fmt.Sprintf("%s/%d", chargeAuthTokenBasePath, chargeAuthTokenID)
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&request); err != nil {
		return nil, err
	}
	return doPatch[ChargeAuthToken](ctx, c, path, &requestBody)
}

// ListChargeAuthTokens to retrieve your charge auth tokens.
func (c *clientImpl) ListChargeAuthTokens(
	ctx context.Context,
	request *ListChargeAuthTokensRequest,
) (*ListChargeAuthTokensResponse, error) {
	query := url.Values{}
	request.PageFilters.Apply(query)
	if request.TeamID != nil {
		query.Set("teamId", strconv.Itoa(int(*request.TeamID)))
	}
	return doGet[ListChargeAuthTokensResponse](ctx, c, chargeAuthTokenBasePath, query)
}

// GetChargeAuthToken to retrieve a single charge auth token.
func (c *clientImpl) GetChargeAuthToken(ctx context.Context, chargeAuthTokenID int64) (*ChargeAuthToken, error) {
	path := fmt.Sprintf("%s/%d", chargeAuthTokenBasePath, chargeAuthTokenID)
	return doGet[ChargeAuthToken](ctx, c, path, nil)
}

// DeleteChargeAuthToken to delete a charge auth token.
func (c *clientImpl) DeleteChargeAuthToken(ctx context.Context, chargeAuthTokenID int64) error {
	path := fmt.Sprintf("%s/%d", chargeAuthTokenBasePath, chargeAuthTokenID)
	return doDelete(ctx, c, path)
}
