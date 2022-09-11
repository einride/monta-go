package monta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// RefreshTokenRequest is the request input to the [Client.RefreshToken] method.
type RefreshTokenRequest struct {
	// The refresh token.
	RefreshToken string `json:"refreshToken"`
}

// RefreshToken creates an authentication token.
func (c *Client) RefreshToken(ctx context.Context, request *RefreshTokenRequest) (_ *Token, err error) {
	const method, path = http.MethodPost, "/v1/auth/refresh"
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s %s: %w", method, path, err)
		}
	}()
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(request); err != nil {
		return nil, err
	}
	httpRequest, err := http.NewRequestWithContext(ctx, method, apiHost+path, &requestBody)
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("content-type", "application/json")
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
	var token Token
	if err := json.NewDecoder(httpResponse.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}
