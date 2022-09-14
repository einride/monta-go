package monta

// WalletTransactionState is a wallet transaction state.
type WalletTransactionState string

// Known [WalletTransactionState] values.
const (
	WalletTransactionStateComplete WalletTransactionState = "complete"
	WalletTransactionStateReserved WalletTransactionState = "reserved"
	WalletTransactionStatePending  WalletTransactionState = "pending"
	WalletTransactionStateFailed   WalletTransactionState = "failed"
)
