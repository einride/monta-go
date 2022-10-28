package monta

import (
	"context"
	"net/url"
	"time"
)

// ListWalletTransactionsRequest is the request input to the [Client.ListWalletTransactions] method.
type ListWalletTransactionsRequest struct {
	PageFilters
	// FromDate allows to filter to retrieve transactions where [WalletTransaction.CreatedAt] >= FromDate.
	FromDate time.Time
	// ToDate allows to filter to retrieve transactions where [WalletTransaction.CreatedAt] <= ToDate.
	ToDate time.Time
}

// ListWalletTransactionsResponse is the response output from the [Client.ListWalletTransactions] method.
type ListWalletTransactionsResponse struct {
	// WalletTransactions in the current page.
	WalletTransactions []*WalletTransaction `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListWalletTransactions to retrieve your wallet transactions.
func (c *Client) ListWalletTransactions(
	ctx context.Context,
	request *ListWalletTransactionsRequest,
) (*ListWalletTransactionsResponse, error) {
	path := "/v1/wallet-transactions"
	query := url.Values{}
	request.PageFilters.Apply(query)
	if !request.FromDate.IsZero() {
		query.Set("fromDate", request.FromDate.UTC().Format(time.RFC3339))
	}
	if !request.ToDate.IsZero() {
		query.Set("toDate", request.ToDate.UTC().Format(time.RFC3339))
	}
	return doGet[ListWalletTransactionsResponse](ctx, c, path, query)
}
