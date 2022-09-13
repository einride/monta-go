package monta

// ChargeState is the state of a charging transaction.
type ChargeState string

// Known [ChargeState] values.
const (
	ChargeStateReserved  ChargeState = "reserved"
	ChargeStateStarting  ChargeState = "starting"
	ChargeStateCharging  ChargeState = "charging"
	ChargeStateStopping  ChargeState = "stopping"
	ChargeStatePaused    ChargeState = "paused"
	ChargeStateScheduled ChargeState = "scheduled"
	ChargeStateStopped   ChargeState = "stopped"
	ChargeStateCompleted ChargeState = "completed"
)
