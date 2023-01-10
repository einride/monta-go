package monta

import (
	"context"
	"net/url"
	"strconv"
	"time"
)

// ListChargesRequest is the request input to the [Client.ListCharges] method.
type ListChargesRequest struct {
	PageFilters

	// Filter to retrieve charges with specified chargePointId.
	ChargePointID *int64

	// Filter to retrieve charges with specified teamId.
	TeamID *int64

	// Filter to retrieve charges with specified siteId.
	SiteID *int64

	// Filter to retrieve charges by state.
	State *ChargeState

	// Filter to retrieve charges by the charge authentication type, must be combined with chargeAuthId.
	ChargeAuthType *ChargeAuthType

	// Filter to retrieve charges by the charge authentication ID, must be combined with chargeAuthType
	// Note: for type vehicleId, chargeAuthId must not include the VID: prefix.
	ChargeAuthID *string

	// Filter to retrieve charges where createdAt >= fromDate.
	FromDate *time.Time
}

// ListChargesResponse is the response output from the [Client.ListCharges] method.
type ListChargesResponse struct {
	// Charges in the current page.
	Charges []*Charge `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListCharges to retrieve your charge points.
func (c *Client) ListCharges(ctx context.Context, request *ListChargesRequest) (*ListChargesResponse, error) {
	path := "/v1/charges"
	query := url.Values{}
	request.PageFilters.Apply(query)
	if request.ChargePointID != nil {
		query.Set("chargePointId", strconv.Itoa(int(*request.ChargePointID)))
	}
	if request.TeamID != nil {
		query.Set("teamId", strconv.Itoa(int(*request.TeamID)))
	}
	if request.SiteID != nil {
		query.Set("siteId", strconv.Itoa(int(*request.SiteID)))
	}
	if request.State != nil {
		query.Set("state", string(*request.State))
	}
	if request.ChargeAuthType != nil {
		query.Set("chargeAuthType", string(*request.ChargeAuthType))
	}
	if request.ChargeAuthID != nil {
		query.Set("chargeAuthId", *request.ChargeAuthID)
	}
	if request.FromDate != nil {
		query.Set("fromDate", request.FromDate.UTC().Format(time.RFC3339))
	}
	return doGet[ListChargesResponse](ctx, c, path, query)
}
