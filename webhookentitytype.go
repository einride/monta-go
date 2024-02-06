package monta

// WebhookEtityType represents the type of a webhook entity.
type WebhookEntityType string

// Known [WebhookEntityType] values.
const (
	WebhookEntityTypeCharge            WebhookEntityType = "charge"
	WebhookEntityTypeChargePoint       WebhookEntityType = "charge-point"
	WebhookEntityTypeSite              WebhookEntityType = "site"
	WebhookEntityTypeTeamMember        WebhookEntityType = "team-member"
	WebhookEntityPluralTypeTeam        WebhookEntityType = "team"
	WebhookEntityTypeInstallerJob      WebhookEntityType = "installer-job"
	WebhookEntityTypeWalletTransaction WebhookEntityType = "wallet-transaction"
	WebhookEntityTypePriceGroup        WebhookEntityType = "price-group"
	WebhookEntityTypePlan              WebhookEntityType = "plan"
	WebhookEntityTypeSubscription      WebhookEntityType = "subscription"
)
