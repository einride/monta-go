package monta

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const apiHost = "https://partner-api.monta.com/api"

type Client interface {
	GetMe(ctx context.Context) (*Me, error)

	// Tokens
	CreateToken(ctx context.Context, request *CreateTokenRequest) (*Token, error)
	RefreshToken(ctx context.Context, request *RefreshTokenRequest) (*Token, error)

	// Charge Auth Tokens
	ListChargeAuthTokens(ctx context.Context, request *ListChargeAuthTokensRequest) (*ListChargeAuthTokensResponse, error)
	GetChargeAuthToken(ctx context.Context, chargeAuthTokenID int64) (*ChargeAuthToken, error)
	CreateChargeAuthToken(ctx context.Context, request CreateChargeAuthTokenRequest) (*ChargeAuthToken, error)
	DeleteChargeAuthToken(ctx context.Context, chargeAuthTokenID int64) error
	PatchChargeAuthToken(
		ctx context.Context,
		chargeAuthTokenID int64,
		request PatchChargeAuthTokenRequest,
	) (*ChargeAuthToken, error)

	// Charge Points
	ListChargePoints(ctx context.Context, request *ListChargePointsRequest) (*ListChargePointsResponse, error)
	GetChargePoint(ctx context.Context, chargePointID int64) (*ChargePoint, error)

	// Charge Point Models
	GetChargePointModel(ctx context.Context, chargePointModelID int64) (*ChargePointModel, error)

	// Charge Point Integrations
	GetChargePointIntegration(ctx context.Context, chargePointIntegrationID int64) (*ChargePointIntegration, error)
	ListChargePointIntegrations(
		ctx context.Context,
		request *ListChargePointIntegrationsRequest,
	) (*ListChargePointIntegrationsResponse, error)

	// Charges
	ListCharges(ctx context.Context, request *ListChargesRequest) (*ListChargesResponse, error)
	GetCharge(ctx context.Context, chargeID int64) (*Charge, error)
	StartCharge(ctx context.Context, request *StartChargeRequest) (*StartChargeResponse, error)
	StopCharge(ctx context.Context, chargeID int64) (*Charge, error)
	RestartCharge(ctx context.Context, chargeID int64) (*Charge, error)

	// Sites
	ListSites(ctx context.Context, request *ListSitesRequest) (*ListSitesResponse, error)
	GetSite(ctx context.Context, siteID int64) (*Site, error)

	// Wallet Transactions
	ListWalletTransactions(
		ctx context.Context,
		request *ListWalletTransactionsRequest,
	) (*ListWalletTransactionsResponse, error)
	GetWalletTransaction(ctx context.Context, transactionID int64) (*WalletTransaction, error)

	// Teams
	ListTeams(
		ctx context.Context,
		request *ListTeamsRequest,
	) (*ListTeamsResponse, error)

	// Webhooks
	GetWebhookConfig(ctx context.Context) (*WebhookConfig, error)
	UpdateWebhookConfig(
		ctx context.Context,
		request *WebhookConfig,
	) (*WebhookConfig, error)
	DeleteWebhookConfig(ctx context.Context) error
	ListWebhookEntries(ctx context.Context, request *ListWebhookEntriesRequest) (*ListWebhookEntriesResponse, error)

	// Price Groups
	ListPriceGroups(ctx context.Context, request *ListPriceGroupsRequest) (*ListPriceGroupsResponse, error)
	CreatePriceGroup(ctx context.Context, request CreateOrUpdatePriceGroupRequest) (*PriceGroup, error)
	GetPriceGroup(ctx context.Context, priceGroupID int64) (*PriceGroup, error)
	UpdatePriceGroup(ctx context.Context, priceGroupID int64, request CreateOrUpdatePriceGroupRequest) (*PriceGroup, error)
	DeletePriceGroup(ctx context.Context, priceGroupID int64) error
	ApplyPriceGroup(ctx context.Context, priceGroupID int64, request ApplyPriceGroupRequest) (*PriceGroup, error)
	SetDefaultPriceGroup(ctx context.Context, priceGroupID int64) error
}

// clientImpl to the Monta Partner API.
type clientImpl struct {
	config         clientConfig
	httpClient     *http.Client
	tokenSemaphore chan struct{}
	token          *Token
}

// Make sure we are implementing all methods.
var _ Client = &clientImpl{}

// ClientOption for configuring a [Client].
type ClientOption func(*clientConfig)

// NewClient creates a new [Client] with the provided [ClientConfig].
func NewClient(options ...ClientOption) Client {
	client := &clientImpl{
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

func (c *clientImpl) setAuthorization(ctx context.Context, request *http.Request) error {
	token, err := c.getToken(ctx)
	if err != nil {
		return err
	}
	request.Header.Set("authorization", "Bearer "+token.AccessToken)
	return nil
}

func (c *clientImpl) getToken(ctx context.Context) (_ *Token, err error) {
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
	if c.token != nil {
		// Add safety margin to now() to prevent token expiration before the request is made
		deadline := time.Now().Add(time.Minute)
		if c.token.AccessTokenExpirationTime.After(deadline) {
			return c.token, nil
		}
		if c.token.RefreshTokenExpirationTime.After(deadline) {
			refreshedToken, err := c.RefreshToken(ctx, &RefreshTokenRequest{
				RefreshToken: c.token.RefreshToken,
			})
			if err != nil {
				return nil, err
			}
			c.token = refreshedToken
			return refreshedToken, nil
		}
	}
	if c.config.clientID == "" || c.config.clientSecret == "" {
		return nil, fmt.Errorf("unable to create token - missing client ID and client secret")
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

// Template method to execute GET requests towards monta.
func doGet[T any](ctx context.Context, client *clientImpl, path string, query url.Values) (*T, error) {
	return execute[T](ctx, client, http.MethodGet, path, query, nil, nil)
}

// Template method to execute POST requests towards monta.
func doPost[T any](ctx context.Context, client *clientImpl, path string, body io.Reader) (*T, error) {
	return execute[T](ctx, client, http.MethodPost, path, nil, &http.Header{"content-type": {"application/json"}}, body)
}

// Template method to execute PATCH requests towards monta.
func doPatch[T any](ctx context.Context, client *clientImpl, path string, body io.Reader) (*T, error) {
	return execute[T](ctx, client, http.MethodPatch, path, nil, nil, body)
}

// Template method to execute DELETE requests towards monta.
func doDelete(ctx context.Context, client *clientImpl, path string) error {
	_, err := execute[any](ctx, client, http.MethodDelete, path, nil, nil, nil)
	return err
}

// Template method to execute PUT requests towards monta.
func doPut[T any](ctx context.Context, client *clientImpl, path string, body io.Reader) (*T, error) {
	return execute[T](ctx, client, http.MethodPut, path, nil, &http.Header{"content-type": {"application/json"}}, body)
}

// Template method to execute PUT requests without a body towards monta.
func doPutNoBody(ctx context.Context, client *clientImpl, path string) error {
	_, err := doPut[struct{}](ctx, client, path, http.NoBody)
	return err
}

// Template method to execute requests towards monta.
func execute[T any](
	ctx context.Context,
	client *clientImpl,
	method string,
	path string,
	query url.Values,
	headers *http.Header,
	body io.Reader,
) (_ *T, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s %s: %w", method, path, err)
		}
	}()
	requestURL, err := url.Parse(apiHost + path)
	if err != nil {
		return nil, err
	}
	if len(query) > 0 {
		requestURL.RawQuery = query.Encode()
	}
	httpRequest, err := http.NewRequestWithContext(ctx, method, requestURL.String(), body)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		httpRequest.Header = *headers
	}
	if err := client.setAuthorization(ctx, httpRequest); err != nil {
		return nil, err
	}
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = httpResponse.Body.Close()
	}()
	if httpResponse.StatusCode != http.StatusOK && httpResponse.StatusCode != http.StatusCreated {
		return nil, newStatusError(httpResponse)
	}
	respBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	if len(respBody) == 0 {
		return nil, nil
	}
	var response T
	return &response, json.Unmarshal(respBody, &response)
}
