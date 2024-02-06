package monta

// WebhookEntityPluralType represents the type of a webhook entity in plural form.
type WebhookEntityPluralType string

// Known [WebhookEntityPluralType] values.
const (
	WebhookEntityPluralTypeAll                WebhookEntityPluralType = "*"
	WebhookEntityPluralTypeCharges            WebhookEntityPluralType = "charges"
	WebhookEntityPluralTypeChargePoints       WebhookEntityPluralType = "charge-points"
	WebhookEntityPluralTypeSites              WebhookEntityPluralType = "sites"
	WebhookEntityPluralTypeTeams              WebhookEntityPluralType = "teams"
	WebhookEntityPluralTypeTeamMembers        WebhookEntityPluralType = "team-members"
	WebhookEntityPluralTypeInstallerJobs      WebhookEntityPluralType = "installer-jobs"
	WebhookEntityPluralTypeWalletTransactions WebhookEntityPluralType = "wallet-transactions"
	WebhookEntityPluralTypePriceGroups        WebhookEntityPluralType = "price-groups"
	WebhookEntityPluralTypePlans              WebhookEntityPluralType = "plans"
	WebhookEntityPluralTypeSubscriptions      WebhookEntityPluralType = "subscriptions"
)
