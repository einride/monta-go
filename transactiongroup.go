package monta

// TransactionGroup is a wallet transaction group.
type TransactionGroup string

// Known [TransactionGroup] values.
const (
	TransactionGroupDeposit  TransactionGroup = "deposit"
	TransactionGroupWithdraw TransactionGroup = "withdraw"
	TransactionGroupCharge   TransactionGroup = "charge"
	TransactionGroupOther    TransactionGroup = "other"
)
