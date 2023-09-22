package monta

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// ListWalletTransactionsRequest is the request input to the [Client.ListWalletTransactions] method.
type ListWalletTransactionsRequest struct {
	PageFilters
	// FromDate allows to filter to retrieve transactions where [WalletTransaction.CreatedAt] >= FromDate.
	FromDate *time.Time
	// ToDate allows to filter to retrieve transactions where [WalletTransaction.CreatedAt] <= ToDate.
	ToDate *time.Time
	TeamID *int64
}

// ListWalletTransactionsResponse is the response output from the [Client.ListWalletTransactions] method.
type ListWalletTransactionsResponse struct {
	// WalletTransactions in the current page.
	WalletTransactions []*WalletTransaction `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListWalletTransactions to retrieve your wallet transactions.
func (c *clientImpl) ListWalletTransactions(
	ctx context.Context,
	request *ListWalletTransactionsRequest,
) (*ListWalletTransactionsResponse, error) {
	path := "/v1/wallet-transactions"
	query := url.Values{}
	request.PageFilters.Apply(query)
	if request.FromDate != nil {
		query.Set("fromDate", request.FromDate.UTC().Format(time.RFC3339))
	}
	if request.ToDate != nil {
		query.Set("toDate", request.ToDate.UTC().Format(time.RFC3339))
	}
	if request.TeamID != nil {
		query.Set("teamId", strconv.Itoa(int(*request.TeamID)))
	}
	return doGet[ListWalletTransactionsResponse](ctx, c, path, query)
}

// GetWalletTransaction to retrieve a single wallet transaction.
func (c *clientImpl) GetWalletTransaction(ctx context.Context, transactionID int64) (*WalletTransaction, error) {
	path := fmt.Sprintf("/v1/wallet-transactions/%d", transactionID)
	return doGet[WalletTransaction](ctx, c, path, nil)
}
