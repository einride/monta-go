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

// ListChargesRequest is the request input to the [Client.ListCharges] method.
type ListChargesRequest struct {
	PageFilters

	// Filter to retrieve charges with specified chargePointId.
	ChargePointID *int64

	// Filter to retrieve charges with specified teamId.
	TeamID *int64

	// Filter to retrieve charges with specified siteId.
	SiteID *int64

	// Filter to retrieve charges with specified operatorId.
	OperatorID *int64

	// Filter to retrieve charges by state.
	State *ChargeState

	// Filter to retrieve charges by the charge authentication type, must be combined with chargeAuthId.
	ChargeAuthType *ChargeAuthType

	// Filter to retrieve charges by the charge authentication ID, must be combined with chargeAuthType
	// Note: for type vehicleId, chargeAuthId must not include the VID: prefix.
	ChargeAuthID *string

	// Filter to retrieve charges where createdAt >= fromDate.
	FromDate *time.Time

	// Filter to retrieve charges where createdAt <= toDate.
	ToDate *time.Time
}

// ListChargesResponse is the response output from the [Client.ListCharges] method.
type ListChargesResponse struct {
	// Charges in the current page.
	Charges []*Charge `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// StartChargeRequest is the request input to the [Client.StartCharge] method.
type StartChargeRequest struct {
	// PayingTeamID is the ID of the team that will be paying for the charge.
	PayingTeamID int64 `json:"payingTeamId"`
	// ChargePointID is the ID of the charge point used for this charge.
	ChargePointID int64 `json:"chargePointId"`
	// ReserveCharge determines whether the charge point will be reserved or start the charge directly.
	ReserveCharge bool `json:"reserveCharge"`
}

// StartChargeResponse is the response output from the [Client.StartCharge] method.
type StartChargeResponse struct {
	// Charge that started.
	Charge Charge `json:"charge"`
}

// ListCharges to retrieve your charges.
func (c *clientImpl) ListCharges(ctx context.Context, request *ListChargesRequest) (*ListChargesResponse, error) {
	path := "/v1/charges"
	query := url.Values{}
	request.Apply(query)
	if request.ChargePointID != nil {
		query.Set("chargePointId", strconv.Itoa(int(*request.ChargePointID)))
	}
	if request.TeamID != nil {
		query.Set("teamId", strconv.Itoa(int(*request.TeamID)))
	}
	if request.SiteID != nil {
		query.Set("siteId", strconv.Itoa(int(*request.SiteID)))
	}
	if request.OperatorID != nil {
		query.Set("operatorId", strconv.Itoa(int(*request.OperatorID)))
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
	if request.ToDate != nil {
		query.Set("toDate", request.ToDate.UTC().Format(time.RFC3339))
	}
	return doGet[ListChargesResponse](ctx, c, path, query)
}

// GetCharge to retrieve a single charge.
func (c *clientImpl) GetCharge(ctx context.Context, chargeID int64) (*Charge, error) {
	path := fmt.Sprintf("/v1/charges/%d", chargeID)
	return doGet[Charge](ctx, c, path, nil)
}

// StartCharge starts a charge.
func (c *clientImpl) StartCharge(ctx context.Context, request *StartChargeRequest) (*StartChargeResponse, error) {
	path := "/v1/charges"
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&request); err != nil {
		return nil, err
	}
	return doPost[StartChargeResponse](ctx, c, path, &requestBody)
}

// StopCharge stops a charge.
func (c *clientImpl) StopCharge(ctx context.Context, chargeID int64) (*Charge, error) {
	path := fmt.Sprintf("/v1/charges/%d/stop", chargeID)
	return doGet[Charge](ctx, c, path, nil)
}

// RestartCharge restarts or starts a reserved charge.
func (c *clientImpl) RestartCharge(ctx context.Context, chargeID int64) (*Charge, error) {
	path := fmt.Sprintf("/v1/charges/%d/restart", chargeID)
	return doGet[Charge](ctx, c, path, nil)
}
