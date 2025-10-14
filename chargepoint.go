package monta

import "time"

// ChargePoint is a charging point.
type ChargePoint struct {
	// ID of this charge point.
	ID int64 `json:"id"`

	// ID of the site
	SiteID *int64 `json:"siteId"`

	// ID of the team.
	TeamID int64 `json:"teamId"`

	// Serial number of this charge point.
	SerialNumber *string `json:"serialNumber"`

	// Name of the site.
	Name *string `json:"name"`

	// Indicates if this charge point is public or private.
	Visibility Visibility `json:"visibility"`

	// Max KW available at this charge point.
	MaxKW *float64 `json:"maxKW"`

	// Type of charge point (AC / DC).
	Type *ChargePointType `json:"type"`

	// A note you have entered for this charge point, e.g. via our Portal.
	Note *string `json:"note"`

	// State of the charge point.
	State *ChargePointState `json:"state"`

	// Last meter reading (KWH) for this charge point.
	LastMeterReadingKwh *float64 `json:"lastMeterReadingKwh"`

	// Indicates if a cable is plugged in (true)
	CablePluggedIn bool `json:"cablePluggedIn"`

	// Brand name for this charge point.
	BrandName *string `json:"brandName"`

	// External Id of this entity, managed by you.
	PartnerExternalID *string `json:"partnerExternalId"`

	// Location of the charge point.
	Location Location `json:"location"`

	// List of supported connector types at this charge point.
	Connectors []Connector `json:"connectors"`

	// DeepLinks to the charge point.
	DeepLinks ChargePointDeepLinks `json:"deeplinks"`

	// When the charge point was created.
	CreatedAt time.Time `json:"createdAt"`

	// When the charge point was last updated.
	UpdatedAt time.Time `json:"updatedAt"`

	// Operator of this charge point
	Operator *Operator `json:"operator"`
}

// ChargePointDeepLinks contains deep-links to a charge point.
type ChargePointDeepLinks struct {
	// Follow this link to open the Monta App with this charge point.
	App string `json:"app"`

	// Follow this link to open the Monta Web App with this charge point.
	Web string `json:"web"`
}
