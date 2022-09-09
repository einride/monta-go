package monta

// Scope is an authorization scope.
type Scope string

// Known [Scope] values.
const (
	ScopeAll                Scope = "all"
	ScopeChargePoints       Scope = "charge-points"
	ScopeMap                Scope = "map"
	ScopeChargeTransactions Scope = "charge-transactions"
	ScopeWalletTransactions Scope = "wallet-transactions"
	ScopeControlCharging    Scope = "control-charging"
)
