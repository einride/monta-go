package monta

import "time"

// ChargeAuthTokenType is the types charge authentications.
type ChargeAuthTokenType string

// Known [ChargeAuthTokenType] values.
const (
	ChargeAuthTokenTypeVehicleID ChargeAuthTokenType = "vehicleId"
	ChargeAuthTokenTypeRFID      ChargeAuthTokenType = "rfid"
)

// ChargeAuthToken used to authenticate the charge.
type ChargeAuthToken struct {
	// The id of the charge auth token.
	ID int64 `json:"id"`
	// The identifier of the charge auth token, Note: without prefix e.g VID:.
	Identifier string `json:"identifier"`
	// The method type used for this charge auth token.
	Type ChargeAuthTokenType `json:"type"`
	// Id of the team that the charge auth token belongs.
	TeamID int64 `json:"teamId"`
	// Name of the charge auth token.
	Name *string `json:"name"`
	// External Id of this entity, managed by you.
	PartnerExternalID *string `json:"partnerExternalId"`
	// Blocked date of this charge auth token.
	BlockedAt *time.Time `json:"blockedAt"`
	// Creation date of this charge auth token.
	CreatedAt time.Time `json:"createdAt"`
	// Update date of this charge auth token.
	UpdatedAt time.Time `json:"updatedAt"`
}
