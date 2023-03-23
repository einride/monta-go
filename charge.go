package monta

import "time"

// SoCSource source of the SoC.
type SoCSource string

// Known [SoCSource] values.
const (
	SoCSourceChargePoint   SoCSource = "charge-point"
	SoCSourceChargeVehicle SoCSource = "vehicle"
)

// Charge is a charging transaction.
type Charge struct {
	// ID of the charge.
	ID int64 `json:"id"`

	// ID of the charge point related to this charge.
	ChargePointID int64 `json:"chargePointId"`

	// When the charge was created.
	CreatedAt time.Time `json:"createdAt"`

	// When the charge was last updated.
	UpdatedAt time.Time `json:"updatedAt"`

	// Date when cable was plugged in.
	CablePluggedInAt *time.Time `json:"cablePluggedInAt"`

	// Date when charge started.
	StartedAt *time.Time `json:"startedAt"`

	// Date when charge stopped.
	StoppedAt *time.Time `json:"stoppedAt"`

	// Date when EV was fully charged.
	FullyChargedAt *time.Time `json:"fullyChargedAt"`

	// Date when charge failed.
	FailedAt *time.Time `json:"failedAt"`

	// Date when charge timed out.
	TimeoutAt *time.Time `json:"timeoutAt"`

	// State of the charge.
	State ChargeState `json:"state"`

	// Consumed Kwh.
	ConsumedKWh *float64 `json:"consumedKwh"`

	// List of consumed Kwh split by hour.
	KwhPerHour []KwhPerHour `json:"kwhPerHour"`

	// Kwh of the meter before charging started.
	StartMeterKWh *float64 `json:"startMeterKwh"`

	// Kwh of the meter after charging stopped.
	EndMeterKWh *float64 `json:"endMeterKwh"`

	// Price for this charge.
	Price *float64 `json:"price"`

	// Configured price limit for this charge.
	PriceLimit *float64 `json:"priceLimit"`

	// Average price per Kwh.
	AveragePricePerKWh *float64 `json:"averagePricePerKwh"`

	// Average CO2 consumption per Kwh.
	AverageCo2PerKWh *float64 `json:"averageCo2PerKwh"`

	// Average percentage of renewable energy per Kwh.
	AverageRenewablePerKWh *float64 `json:"averageRenewablePerKwh"`

	// Failure reason for this charge.
	FailureReason *string `json:"failureReason"`

	// Payment method for this charge.
	PaymentMethod *PaymentMethod `json:"paymentMethod"`

	// A note taken for this charge.
	Note *string `json:"note"`

	// Configured Kwh limit for this charge.
	KWhLimit *float64 `json:"kwhLimit"`

	// Currency for paying the charge.
	Currency *Currency `json:"currency"`

	// PayingTeam is the team paying for the charge.
	PayingTeam *Team `json:"payingTeam"`

	// ChargeAuth is the method used to authenticate the charge.
	ChargeAuth *ChargeAuth `json:"chargeAuth"`

	// Information about the state of charge.
	SoC *SoC `json:"soc"`

	// Configured SoC limit for this charge.
	SoCLimit *float64 `json:"socLimit"`
}

// Sum of kwh for a given hour.
type KwhPerHour struct {
	// Hour for the sum of the kwh.
	Time time.Time `json:"time"`

	// Sum of kwh for this hour.
	Value float64 `json:"value"`
}

// Information about the state of charge if available.
type SoC struct {
	// Value of SoC in %
	Percentage *float64 `json:"percentage"`

	// Source of this value, eg vehicle or charge-point.
	Source SoCSource `json:"source"`
}
