package monta

// OperatorRole enumerates filters for the operator's relationship to listed charges.
type OperatorRole string

// Known [OperatorRole] values.
const (
	// OperatorRoleOwner returns charges from charge points owned by the operator.
	OperatorRoleOwner OperatorRole = "owner"
	// OperatorRolePayer returns charges paid by the operator.
	OperatorRolePayer OperatorRole = "payer"
)
