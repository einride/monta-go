package monta

// ChargeAuthType is the types charge authentications.
type ChargeAuthType string

// Known [ChargeAuthType] values.
const (
	ChargeAuthTypeVehicleID ChargeAuthType = "vehicleId"
	ChargeAuthTypeRFID      ChargeAuthType = "rfid"
	ChargeAuthTypeApp       ChargeAuthType = "app"
)

// ChargeAuth method used to authenticate the charge.
type ChargeAuth struct {
	// The method type used to authenticate a charge.
	Type ChargeAuthType `json:"type"`

	// The id of the chosen authentication method.
	ID string `json:"id"`
}
