package monta

// ChargeState is the state of a charging transaction.
type ChargeState string

// Known [ChargeState] values.
const (
	ChargeStatePaying    ChargeState = "paying"
	ChargeStateReserved  ChargeState = "reserved"
	ChargeStateStarting  ChargeState = "starting"
	ChargeStateCharging  ChargeState = "charging"
	ChargeStateStopping  ChargeState = "stopping"
	ChargeStatePaused    ChargeState = "paused"
	ChargeStateScheduled ChargeState = "scheduled"
	ChargeStateStopped   ChargeState = "stopped"
	ChargeStateCompleted ChargeState = "completed"
	ChargeStateReleasing ChargeState = "releasing"
	ChargeStateReleased  ChargeState = "released"
	ChargeStateOther     ChargeState = "other"
)
