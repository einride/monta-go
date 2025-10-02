package monta

import "time"

type ChargePointIntegration struct {
	// The id of this charge point integration.
	ID int64 `json:"id"`

	// Enumerate the possible states for a charge point integration
	State ChargePointIntegrationState `json:"state"`

	// The serial number of this charge point integration
	SerialNumber string `json:"serialNumber"`

	// The identity number of this charge point integration
	ChargePointIdentity string `json:"chargePointIdentity"`

	// Identifier of the integration type.
	IntegrationTypeIdentifier string `json:"integrationTypeIdentifier"`

	// The connector id for this charge point integration
	ConnectorID *int64 `json:"connectorId"`

	// The charge point for this charge point integration
	ChargePoint ChargePoint `json:"chargePoint"`

	// The Date this charge point integration was set to active
	ActiveAt *time.Time `json:"activeAt"`

	// The Creation date of this charge point integration
	CreatedAt time.Time `json:"createdAt"`

	// The Date this charge point integration was last updated
	UpdatedAt time.Time `json:"updatedAt"`

	// The Date this charge point integration was deleted
	DeletedAt *time.Time `json:"deletedAt"`
}

type ChargePointIntegrationState string

const (
	ChargePointIntegrationStatePending      ChargePointIntegrationState = "pending"
	ChargePointIntegrationStateConnected    ChargePointIntegrationState = "connected"
	ChargePointIntegrationStateDisconnected ChargePointIntegrationState = "disconnected"
	ChargePointIntegrationStateError        ChargePointIntegrationState = "error"
	ChargePointIntegrationStateUnknown      ChargePointIntegrationState = "unknown"
)
