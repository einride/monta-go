package monta

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const apiHost = "https://partner-api.monta.app/api"

// Client to the Monta Partner API.
type Client struct {
	config         clientConfig
	httpClient     *http.Client
	tokenSemaphore chan struct{}
	token          *Token
}

// ClientOption for configuring a [Client].
type ClientOption func(*clientConfig)

// NewClient creates a new [Client] with the provided [ClientConfig].
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		httpClient:     http.DefaultClient,
		tokenSemaphore: make(chan struct{}, 1),
	}
	for _, option := range options {
		option(&client.config)
	}
	client.tokenSemaphore <- struct{}{}
	client.token = client.config.token
	return client
}

type clientConfig struct {
	clientID     string
	clientSecret string
	token        *Token
}

// WithClientIDAndSecret configures authentication using the provided client ID and secret.
func WithClientIDAndSecret(clientID, clientSecret string) ClientOption {
	return func(config *clientConfig) {
		config.clientID = clientID
		config.clientSecret = clientSecret
	}
}

// WithToken configures authentication using the provided authentication token.
func WithToken(token *Token) ClientOption {
	return func(config *clientConfig) {
		config.token = token
	}
}

func (c *Client) setAuthorization(ctx context.Context, request *http.Request) error {
	token, err := c.getToken(ctx)
	if err != nil {
		return err
	}
	request.Header.Set("authorization", "Bearer "+token.AccessToken)
	return nil
}

func (c *Client) getToken(ctx context.Context) (_ *Token, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("get token: %w", err)
		}
	}()
	select {
	case <-c.tokenSemaphore:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	defer func() {
		c.tokenSemaphore <- struct{}{}
	}()
	cachedToken, ok := c.getCachedToken()
	if ok {
		return cachedToken, nil
	}
	if c.config.clientID == "" || c.config.clientSecret == "" {
		return nil, fmt.Errorf("unable to create token - missing Client ID and Client Secret")
	}
	createdToken, err := c.CreateToken(ctx, &CreateTokenRequest{
		ClientID:     c.config.clientID,
		ClientSecret: c.config.clientSecret,
	})
	if err != nil {
		return nil, err
	}
	c.token = createdToken
	return createdToken, nil
}

func (c *Client) getCachedToken() (*Token, bool) {
	if c.token != nil && c.token.AccessTokenExpirationTime.After(time.Now()) {
		return c.token, true
	}
	return nil, false
}
