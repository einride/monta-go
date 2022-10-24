package monta

import (
	"context"
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
	path := "/v1/charge-points"
	query := url.Values{}
	if request.Page > 0 {
		query.Set("page", strconv.Itoa(request.Page))
	}
	if request.PerPage > 0 {
		query.Set("perPage", strconv.Itoa(request.PerPage))
	}
	if request.SiteID != nil {
		query.Set("siteId", strconv.Itoa(int(*request.SiteID)))
	}
	return doGet[ListChargePointsResponse](ctx, c, path, query)
}
