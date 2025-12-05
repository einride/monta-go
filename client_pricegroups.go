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

const priceGroupBasePath = "/v1/price-groups"

// CreatePriceGroupRequest is the request input to the [Client.CreatePriceGroup] method.
type CreateOrUpdatePriceGroupRequest struct {
	// Team ID.
	TeamID int64 `json:"teamId"`
	// Name of the price group.
	Name string `json:"name"`
	// Type of the price group.
	Type PriceGroupType `json:"type"`
	// The master price.
	MasterPrice Pricing `json:"masterPrice"`
	// All the fees for the price group.
	Fees *[]Pricing `json:"fees"`
}

// ListPriceGroupsRequest is the request input to the [Client.ListPriceGroups] method.
type ListPriceGroupsRequest struct {
	PageFilters
	// Filter to retrieve price groups with specified team ID.
	TeamID *int64
}

// ListPriceGroupsResponse is the response output from the [Client.ListPriceGroups] method.
type ListPriceGroupsResponse struct {
	// Price groups in the current page.
	PriceGroups []*PriceGroup `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ApplyPriceGroupRequest is the request input to the [Client.ApplyPriceGroup] method.
type ApplyPriceGroupRequest struct {
	// List of charge point IDs.
	SelectedChargePointIDs *[]int64 `json:"selectedChargePointIds"`
	// List of site IDs.
	SelectedSiteIDs *[]int64 `json:"selectedSiteIds"`
	// List of team member IDs.
	SelectedTeamMemberIDs *[]int64 `json:"selectedTeamMemberIds"`
	// Flag to indicate if costs should be reset.
	ResetCost *bool `json:"resetCost"`
}

// A pricing object for a create or update request body.
type PricingRequestBody struct {
	// Name of the pricing. It will be null when it's the master price.
	Description *string `json:"description"`
	// Type of the pricing. minute is used for Minute fee. min is used for the master price.
	Type PricingType `json:"type"`
	// Used by the Minute fee. True means it will stop charging the fee when the charge is complete.
	// False means it will stop charging the fee when the cable is unplugged.
	EndAtFullyCharged bool `json:"endAtFullyCharged"`
	// Used by Spot Price. True means it will add % of VAT on top the price calculations.
	// Note: vat rates differ from country to country.
	VAT bool `json:"vat"`
	// Used by Spot Price. It will multiply the fallback price by this percentage.
	Percentage *float64 `json:"percentage"`
	// The id of the selected Tariff
	TariffID *int64 `json:"tariffId"`
	// The id of the charge pricing tag for this pricing.
	TagID *int64 `json:"tagId"`
	// Used by Charging, Minute and Idle Fees. After how many minutes the fee should start being applied.
	ApplyAfterMinutes *int32 `json:"applyAfterMinutes"`
	// The price of this Fee or Master price.
	Price Price `json:"price"`
	// Used by spot price. The minimum that the raw spot price can be.
	// This will be used in calculations if spot price is lower than this.
	PriceMin *Price `json:"priceMin"`
	// Used by spot price. The maximum that the raw spot price can be.
	// This will be used in calculations if spot price is higher than this.
	PriceMax *Price `json:"priceMax"`
	// Used by Idle fee. The maximum the user will be charged for the idle fee.
	FeePriceMax *Price `json:"feePriceMax"`
	// Used by spot price. Additional absolute money or percentages values to be added on top of the previous calculations.
	Additional []*Additional `json:"additional"`
	// DateTime "from" time to which this pricing should apply from.
	From *time.Time `json:"from"`
	// DateTime "to" time to which this pricing should apply to
	To *time.Time `json:"to"`
}

// ListPriceGroups to retrieve your price groups.
func (c *clientImpl) ListPriceGroups(
	ctx context.Context,
	request *ListPriceGroupsRequest,
) (*ListPriceGroupsResponse, error) {
	query := url.Values{}
	request.Apply(query)
	if request.TeamID != nil {
		query.Set("teamId", strconv.Itoa(int(*request.TeamID)))
	}
	return doGet[ListPriceGroupsResponse](ctx, c, priceGroupBasePath, query)
}

// CreatePriceGroup to create a new price group.
func (c *clientImpl) CreatePriceGroup(
	ctx context.Context,
	request CreateOrUpdatePriceGroupRequest,
) (*PriceGroup, error) {
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&request); err != nil {
		return nil, err
	}
	return doPost[PriceGroup](ctx, c, priceGroupBasePath, &requestBody)
}

// GetPriceGroup to retrieve a single price group.
func (c *clientImpl) GetPriceGroup(ctx context.Context, priceGroupID int64) (*PriceGroup, error) {
	path := fmt.Sprintf("%s/%d", priceGroupBasePath, priceGroupID)
	return doGet[PriceGroup](ctx, c, path, nil)
}

// UpdatePriceGroup to update a price group.
func (c *clientImpl) UpdatePriceGroup(
	ctx context.Context,
	priceGroupID int64,
	request CreateOrUpdatePriceGroupRequest,
) (*PriceGroup, error) {
	path := fmt.Sprintf("%s/%d", priceGroupBasePath, priceGroupID)
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&request); err != nil {
		return nil, err
	}
	return doPut[PriceGroup](ctx, c, path, &requestBody)
}

// DeletePriceGroup to delete a price group.
func (c *clientImpl) DeletePriceGroup(ctx context.Context, priceGroupID int64) error {
	path := fmt.Sprintf("%s/%d", priceGroupBasePath, priceGroupID)
	return doDelete(ctx, c, path)
}

// ApplyPriceGroup to aply a price group to charge points, sites or team members.
func (c *clientImpl) ApplyPriceGroup(
	ctx context.Context,
	priceGroupID int64,
	request ApplyPriceGroupRequest,
) (*PriceGroup, error) {
	path := fmt.Sprintf("%s/%d/apply", priceGroupBasePath, priceGroupID)
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&request); err != nil {
		return nil, err
	}
	return doPut[PriceGroup](ctx, c, path, &requestBody)
}

// SetDefaultPriceGroup to set a price group as default.
func (c *clientImpl) SetDefaultPriceGroup(ctx context.Context, priceGroupID int64) error {
	path := fmt.Sprintf("%s/%d/default", priceGroupBasePath, priceGroupID)
	return doPutNoBody(ctx, c, path)
}
