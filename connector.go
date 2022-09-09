package monta

// Connector is a charge point connector.
type Connector struct {
	// Identifier of connector
	Identifier string `json:"identifier"`
	// Readable name of connector.
	Name string `json:"name"`
}
