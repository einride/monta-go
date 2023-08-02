package monta

// ChargePointState represents the state of a charge point.
type ChargePointState string

// Known [ChargePointState] values.
const (
	ChargePointStateAvailable       ChargePointState = "available"
	ChargePointStateBusy            ChargePointState = "busy"
	ChargePointStateBusyBlocked     ChargePointState = "busy-blocked"
	ChargePointStateBusyCharging    ChargePointState = "busy-charging"
	ChargePointStateBusyNonCharging ChargePointState = "busy-non-charging"
	ChargePointStateBusyNonReleased ChargePointState = "busy-non-released"
	ChargePointStateBusyReserved    ChargePointState = "busy-reserved"
	ChargePointStateBusyScheduled   ChargePointState = "busy-scheduled"
	ChargePointStateError           ChargePointState = "error"
	ChargePointStateDisconnected    ChargePointState = "disconnected"
	ChargePointStatePassive         ChargePointState = "passive"
	ChargePointStateOther           ChargePointState = "other"
)
