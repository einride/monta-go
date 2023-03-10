package monta

import "time"

// Site is a charging site.
type Site struct {
	// ID of the site.
	ID int64 `json:"id"`

	// Name of the site.
	Name string `json:"name"`

	// ChargePointCount is the number of charge points at this site.
	ChargePointCount int64 `json:"chargePointCount"`

	// ActiveChargePointCount is the number of active charge points at this site.
	ActiveChargePointCount int64 `json:"activeChargePointCount"`

	// AvailableChargePointCount is the number of available charge points at this site.
	AvailableChargePointCount int64 `json:"availableChargePointCount"`

	// MaxKW available at this site.
	MaxKW *float64 `json:"maxKW"`

	// Type of charge points at this site.
	Type *ChargePointType `json:"type"`

	// Visibility indicates if this site is public or private.
	Visibility Visibility `json:"visibility"`

	// A Note you have entered for this site, e.g. via our Portal.
	Note *string `json:"note"`

	// Location of the site.
	Location Location `json:"location"`

	// Connectors is a list of supported connector types at this site.
	Connectors []Connector `json:"connectors"`

	// When the charging site was created.
	CreatedAt time.Time `json:"createdAt"`

	// When the charging site was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
}
