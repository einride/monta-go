package monta

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
func (c *Client) GetMe(ctx context.Context) (_ *Me, err error) {
	const method, path = http.MethodGet, "/v1/auth/me"
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s %s: %w", method, path, err)
		}
	}()
	httpRequest, err := http.NewRequestWithContext(ctx, method, apiHost+path, nil)
	if err != nil {
		return nil, err
	}
	if err := c.setAuthorization(ctx, httpRequest); err != nil {
		return nil, err
	}
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = httpResponse.Body.Close()
	}()
	if httpResponse.StatusCode != http.StatusOK {
		return nil, newStatusError(httpResponse)
	}
	var me Me
	if err := json.NewDecoder(httpResponse.Body).Decode(&me); err != nil {
		return nil, err
	}
	return &me, nil
}
