package monta

import "time"

// ChargePointModel is a charge point model.
type ChargePointModel struct {
	// ID is the unique numeric identifier for the model.
	ID int64 `json:"id"`
	// Identifier is the unique string identifier for the model (e.g., "rolec_securicharge_ev_dual").
	Identifier string `json:"identifier"`
	// Name is the display name of the model (e.g., "SecuriCharge Dual").
	Name string `json:"name"`
	// Brand is the manufacturer of the charge point model.
	Brand *Brand `json:"brand"`
	// Features is a list of supported features for this model.
	Features []Feature `json:"features"`
	// CreatedAt is the timestamp of when the model was created.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt is the timestamp of the last update.
	UpdatedAt time.Time `json:"updatedAt"`
}

// Brand represents the manufacturer of a charge point (e.g., "Rolec", "Kempower").
type Brand struct {
	// ID is the unique numeric identifier for the brand.
	ID int64 `json:"id"`
	// Name is the display name of the brand.
	Name string `json:"name"`
}

// Feature represents a specific capability of a charge point model.
type Feature struct {
	// Key is the programmatic identifier for the feature (e.g., "firmware_management").
	Key string `json:"key"`
	// Description is a human-readable explanation of the feature.
	Description string `json:"description"`
	// Requirements describes what is needed for this feature to work.
	Requirements string `json:"requirements"`
	// Enabled indicates whether the feature is active.
	Enabled bool `json:"enabled"`
}
