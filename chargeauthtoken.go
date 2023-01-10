package monta

import "time"

// ChargeAuthToken used to authenticate the charge.
type ChargeAuthToken struct {
	// The id of the charge auth token.
	ID int64 `json:"id"`
	// The identifier of the charge auth token, Note: without prefix e.g VID:.
	Identifier string `json:"identifier"`
	// The method type used for this charge auth token.
	Type ChargeAuthType `json:"type"`
	// Id of the team that the charge auth token belongs.
	TeamID int64 `json:"teamId"`
	// Name of the charge auth token.
	Name *string `json:"name"`
	// Blocked date of this charge auth token.
	BlockedAt *time.Time `json:"blockedAt"`
	// Creation date of this charge auth token.
	CreatedAt time.Time `json:"createdAt"`
	// Update date of this charge auth token.
	UpdatedAt time.Time `json:"updatedAt"`
}
