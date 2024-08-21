package monta

// Operator is a charge point operator.
type Operator struct {
	// ID of the operator.
	ID int64 `json:"id"`

	// Name of operator.
	Name string `json:"name"`

	// Identifier of operator.
	Identifier string `json:"identifier"`

	// VATNumber of the operator.
	VATNumber string `json:"vatNumber"`

	// PartnerID of the operator (owner of this operator).
	PartnerID int64 `json:"partnerId"`
}
