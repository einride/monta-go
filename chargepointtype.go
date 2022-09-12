package monta

// ChargePointType is the type of charge point.
type ChargePointType string

// Known [ChargePointType] values.
const (
	ChargePointTypeAC ChargePointType = "ac"
	ChargePointTypeDC ChargePointType = "dc"
)
