package monta

// Currency represents a currency.
type Currency struct {
	// ID of the currency, e.g. DKK.
	ID string `json:"identifier"`

	// Readable name of currency.
	Name string `json:"name"`

	// Number of decimals for this currency.
	Decimals int32 `json:"decimals"`
}
