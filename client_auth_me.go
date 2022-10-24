package monta

import (
	"context"
)

// Me - the current API consumer.
type Me struct {
	// My Name (e.g. Monta Team A).
	Name string `json:"name"`

	// My OperatorID.
	OperatorID int64 `json:"operatorId"`

	// TeamIDs that are unlocked for API operations. If empty, all teams of this operator are unlocked.
	TeamIDs []int64 `json:"teamIds"`

	// My ClientID.
	ClientID string `json:"clientId"`

	// RateLimit for my account; requests per minute.
	RateLimit int64 `json:"rateLimit"`

	// My Scopes.
	Scopes []Scope `json:"scopes"`
}

// GetMe obtains information about current API consumer [Me].
func (c *Client) GetMe(ctx context.Context) (*Me, error) {
	return doGet[Me](ctx, c, "/v1/auth/me", nil)
}
