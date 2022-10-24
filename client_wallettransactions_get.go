package monta

import (
	"context"
	"fmt"
)

// GetWalletTransaction to retrieve a single wallet transaction.
func (c *Client) GetWalletTransaction(ctx context.Context, transactionID int64) (*WalletTransaction, error) {
	path := fmt.Sprintf("/v1/wallet-transactions/%d", transactionID)
	return doGet[WalletTransaction](ctx, c, path, nil)
}
