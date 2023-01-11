package monta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Token holds authentication tokens for the Monta Partner API.
type Token struct {
	// AccessToken for accessing the Monta Partner API.
	AccessToken string `json:"accessToken"`
	// RefreshToken for refreshing an access token.
	RefreshToken string `json:"refreshToken"`
	// AccessTokenExpirationDate is the expiration time of the access token.
	AccessTokenExpirationTime time.Time `json:"accessTokenExpirationDate"`
	// RefreshTokenExpirationDate is the expiration date of the access token.
	RefreshTokenExpirationTime time.Time `json:"refreshTokenExpirationDate"`
}

// CreateTokenRequest is the request input to the [Client.CreateToken] method.
type CreateTokenRequest struct {
	// The ClientID to use.
	ClientID string `json:"clientId"`
	// The ClientSecret to use.
	ClientSecret string `json:"clientSecret"`
}

// RefreshTokenRequest is the request input to the [Client.RefreshToken] method.
type RefreshTokenRequest struct {
	// The refresh token.
	RefreshToken string `json:"refreshToken"`
}

// CreateToken creates an authentication token.
func (c *Client) CreateToken(ctx context.Context, request *CreateTokenRequest) (_ *Token, err error) {
	const method, path = http.MethodPost, "/v1/auth/token"
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
