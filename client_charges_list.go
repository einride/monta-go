package monta

import (
	"context"
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
func (c *Client) ListCharges(ctx context.Context, request *ListChargesRequest) (*ListChargesResponse, error) {
	path := "/v1/charges"
	query := url.Values{}
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
	return listEntity[ListChargesResponse](ctx, c, path, query)
}
