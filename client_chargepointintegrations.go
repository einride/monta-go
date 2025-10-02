package monta

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// GetChargePointIntegration to retrieve a charge point integration.
func (c *clientImpl) GetChargePointIntegration(
	ctx context.Context,
	chargePointIntegrationID int64,
) (*ChargePointIntegration, error) {
	path := fmt.Sprintf("/v1/charge-point-integrations/%d", chargePointIntegrationID)
	return doGet[ChargePointIntegration](ctx, c, path, nil)
}

// ListChargePointIntegrationsRequest is the request output from the [Client.ListChargePointIntegrations] method.
type ListChargePointIntegrationsRequest struct {
	PageFilters
	// Allows to filter list of charge points integrations by a charge point id
	ChargePointID int64
	// Includes deleted resources in the response
	IncludeDeleted *bool
}

// ListChargePointIntegrationsResponse is the response output from the [Client.ListChargePointIntegrations] method.
type ListChargePointIntegrationsResponse struct {
	// List of charges point integrations that match the criteria.
	ChargePointIntegrations []*ChargePointIntegration `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListChargePoints to retrieve your charge point integrations.
func (c *clientImpl) ListChargePointIntegrations(
	ctx context.Context,
	request *ListChargePointIntegrationsRequest,
) (*ListChargePointIntegrationsResponse, error) {
	path := "/v1/charge-point-integrations"
	query := url.Values{}
	request.PageFilters.Apply(query)
	query.Set("chargePointId", strconv.Itoa(int(request.ChargePointID)))
	if request.IncludeDeleted != nil {
		query.Set("includeDeleted", strconv.FormatBool(*request.IncludeDeleted))
	}
	return doGet[ListChargePointIntegrationsResponse](ctx, c, path, query)
}
