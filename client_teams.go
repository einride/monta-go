package monta

import (
	"context"
	"net/url"
)

// ListTeamsRequest is the request input to the [Client.ListTeams] method.
type ListTeamsRequest struct {
	PageFilters
	// Filter teams by partner external id. To filter only resources without PartnerExternalID use "".
	PartnerExternalID *string
	// If the team can be deleted.
	IncludeDeleted bool
}

// ListTeamsResponse is the response output from the [Client.ListTeams] method.
type ListTeamsResponse struct {
	// Teams in the current page.
	Teams []*Team `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListTeams to retrieve your teams.
func (c *clientImpl) ListTeams(
	ctx context.Context,
	request *ListTeamsRequest,
) (*ListTeamsResponse, error) {
	path := "/v1/teams"
	query := url.Values{}
	request.Apply(query)
	if request.PartnerExternalID != nil {
		query.Set("partnerExternalId", *request.PartnerExternalID)
	}
	if request.IncludeDeleted {
		query.Set("includeDeleted", "true")
	}
	return doGet[ListTeamsResponse](ctx, c, path, query)
}
