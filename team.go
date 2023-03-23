package monta

// Team of Monta users.
type Team struct {
	// ID of the team.
	ID int64 `json:"id"`

	// Public name of the team.
	PublicName string `json:"publicName"`

	// External Id of this entity, managed by you.
	PartnerExternalID *string `json:"partnerExternalId"`
}
