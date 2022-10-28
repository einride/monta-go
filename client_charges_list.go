package monta

import (
	"context"
	"net/url"
	"strconv"
)

// ListChargesRequest is the request input to the [Client.ListCharges] method.
type ListChargesRequest struct {
	PageFilters
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
func (c *Client) ListCharges(ctx context.Context, request *ListChargesRequest) (*ListChargesResponse, error) {
	path := "/v1/charges"
	query := url.Values{}
	request.PageFilters.Apply(query)
	if request.TeamID != nil {
		query.Set("teamId", strconv.Itoa(int(*request.TeamID)))
	}
	if request.ChargePointID != nil {
		query.Set("chargePointId", strconv.Itoa(int(*request.ChargePointID)))
	}
	return doGet[ListChargesResponse](ctx, c, path, query)
}
