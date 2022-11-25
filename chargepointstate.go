package monta

// ChargePointState represents the state of a charge point.
type ChargePointState string

// Known [ChargePointState] values.
const (
	ChargePointStateAvailable    ChargePointState = "available"
	ChargePointStateBusy         ChargePointState = "busy"
	ChargePointStateBusyCharging ChargePointState = "busy-charging"
	ChargePointStateError        ChargePointState = "error"
	ChargePointStateDisconnected ChargePointState = "disconnected"
	ChargePointStatePassive      ChargePointState = "passive"
	ChargePointStateOther        ChargePointState = "other"
)
