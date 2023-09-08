package monta

import "time"

// Team of Monta users.
type Team struct {
	// ID of the team.
	ID int64 `json:"id"`

	// Name of the team.
	Name string `json:"name"`

	// External Id of the team.
	ExternalID *string `json:"externalId"`

	// External Id of this entity, managed by you.
	PartnerExternalID *string `json:"partnerExternalId"`

	// Code to share with a user to join the team.
	JoinCode string `json:"joinCode"`

	// Company name for the given team.
	CompanyName *string `json:"companyName"`

	// Operator of the team.
	Operator Operator `json:"operator"`

	// Address of the team.
	Address Address `json:"address"`

	// Type of the team.
	Type *string `json:"type"`

	// Operator Id of the team.
	OperatorID int64 `json:"operatorId"`

	// When the team was blocked.
	BlockedAt *time.Time `json:"blockedAt"`

	// When the team was created.
	CreatedAt time.Time `json:"createdAt"`

	// When the team was last updated.
	UpdatedAt time.Time `json:"updatedAt"`

	// When the team was deleted.
	DeletedAt *time.Time `json:"deletedAt"`
}
