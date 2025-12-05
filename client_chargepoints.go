package monta

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ListChargePointsRequest is the request input to the [Client.ListChargePoints] method.
type ListChargePointsRequest struct {
	PageFilters
	// SiteID allows to filter list of charge points by a site id.
	SiteID *int64
	// TeamID allows to filter list of charge points by a team id.
	TeamID *int64
}

// ListChargePointsResponse is the response output from the [Client.ListChargePoints] method.
type ListChargePointsResponse struct {
	// ChargePoints in the current page.
	ChargePoints []*ChargePoint `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListChargePoints to retrieve your charge points.
func (c *clientImpl) ListChargePoints(
	ctx context.Context,
	request *ListChargePointsRequest,
) (*ListChargePointsResponse, error) {
	path := "/v1/charge-points"
	query := url.Values{}
	request.Apply(query)
	if request.SiteID != nil {
		query.Set("siteId", strconv.Itoa(int(*request.SiteID)))
	}
	if request.TeamID != nil {
		query.Set("teamId", strconv.Itoa(int(*request.TeamID)))
	}
	return doGet[ListChargePointsResponse](ctx, c, path, query)
}

// GetChargePoint to retrieve a single charge point.
func (c *clientImpl) GetChargePoint(ctx context.Context, chargePointID int64) (*ChargePoint, error) {
	path := fmt.Sprintf("/v1/charge-points/%d", chargePointID)
	return doGet[ChargePoint](ctx, c, path, nil)
}
