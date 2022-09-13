package monta

// PaymentMethod is a payment method.
type PaymentMethod string

// Known [PaymentMethod] values.
const (
	PaymentMethodFree              PaymentMethod = "free"
	PaymentMethodTeamHasFunds      PaymentMethod = "team-has-funds"
	PaymentMethodTeamHasAutoRefill PaymentMethod = "team-has-auto-refill"
	PaymentMethodSource            PaymentMethod = "source"
	PaymentMethodPayment           PaymentMethod = "payment"
)
