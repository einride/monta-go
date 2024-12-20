package monta

import "time"

// PriceGroup is a price group.
type PriceGroup struct {
	// ID of the price group.
	ID int64 `json:"id"`

	// Name of the price group.
	Name string `json:"name"`

	// Default price group.
	Default bool `json:"default"`

	// Type of the price group.
	Type PriceGroupType `json:"type"`

	// The master price.
	MasterPrice Pricing `json:"masterPrice"`

	// Tariffs of the price group.
	Tariffs []Pricing `json:"tariffs"`

	// Fees of the price group.
	Fees []Pricing `json:"fees"`

	// To how many team members the price group has been applied to.
	TeamMemberCount *int32 `json:"teamMemberCount"`

	// To how many charge points the price group has been applied to.
	ChargePointCount *int32 `json:"chargePointCount"`

	// Objects the price group is applied to.
	AppliedTo *AppliedTo `json:"appliedTo"`

	// When the price group was created.
	CreatedAt time.Time `json:"createdAt"`

	// When the price group was updated.
	UpdatedAt *time.Time `json:"updatedAt"`
}

// Type of price group.
type PriceGroupType string

// Known [PriceGroupType] values.
const (
	PriceGroupTypePublic        PriceGroupType = "public"
	PriceGroupTypeMember        PriceGroupType = "member"
	PriceGroupTypeSponsored     PriceGroupType = "sponsored"
	PriceGroupTypeCost          PriceGroupType = "cost"
	PriceGroupTypeRoaming       PriceGroupType = "roaming"
	PriceGroupTypeReimbursement PriceGroupType = "reimbursement"
	PriceGroupTypeOther         PriceGroupType = "other"
)

// A pricing object.
type Pricing struct {
	// Id of the pricing.
	ID int64 `json:"id"`

	// Name of the pricing. It will be null when it's the master price.
	Description *string `json:"description"`

	// Type of the pricing. minute is used for Minute fee. min is used for the master price.
	Type PricingType `json:"type"`

	// If this is the master price (not a fee).
	Master bool `json:"master"`

	// If it's a dynamic price. It will be true if a tariffId is present.
	DynamicPricing bool `json:"dynamicPricing"`

	// Used by the Minute fee. True means it will stop charging the fee when the charge is complete.
	// False means it will stop charging the fee when the cable is unplugged.
	EndAtFullyCharged bool `json:"endAtFullyCharged"`

	// Used by Spot Price. True means it will add % of VAT on top the price calculations.
	// Note: vat rates differ from country to country.
	VAT bool `json:"vat"`

	// Used by Spot Price. It will multiply the fallback price by this percentage.
	Percentage *float64 `json:"percentage"`

	// The id of the selected Tariff
	TariffID *int64 `json:"tariffId"`

	// When the pricing was last updated.
	UpdatedAt time.Time `json:"updatedAt"`

	// Used by Charging, Minute and Idle Fees. After how many minutes the fee should start being applied.
	ApplyAfterMinutes *int32 `json:"applyAfterMinutes"`

	// The price of this Fee or Master price.
	Price Price `json:"price"`

	// Used by spot price. The minimum that the raw spot price can be.
	// This will be used in calculations if spot price is lower than this.
	PriceMin *Price `json:"priceMin"`

	// Used by spot price. The maximum that the raw spot price can be.
	// This will be used in calculations if spot price is higher than this.
	PriceMax *Price `json:"priceMax"`

	// Used by Idle fee. The maximum the user will be charged for the idle fee.
	FeePriceMax *Price `json:"feePriceMax"`

	// Used by spot price. Additional absolute money or percentages values to be added on top of the previous calculations.
	Additional []*Additional `json:"additional"`

	// DateTime "from" time to which this pricing should apply from.
	From *time.Time `json:"from"`

	// DateTime "to" time to which this pricing should apply to
	To *time.Time `json:"to"`

	// The id of the charge pricing tag for this pricing.
	TagID *int64 `json:"tagId"`
}

// Type of pricing.
type PricingType string

// Know [PricingType] values.
const (
	PricingTypeKwh      PricingType = "kwh"
	PricingTypeMin      PricingType = "min"
	PricingTypeSpot     PricingType = "spot"
	PricingTypeTariff   PricingType = "tariff"
	PricingTypeStarting PricingType = "starting"
	PricingTypeCharging PricingType = "charging"
	PricingTypeIdle     PricingType = "idle"
	PricingTypeMinute   PricingType = "minute"
)

// Price of a fee, tariff or a master price.
type Price struct {
	// The amount of money.
	Amount int64 `json:"amount"`
	// Currency.
	Currency PriceCurrency `json:"currency"`
	// Current user locale.
	Locale string `json:"locale"`
}

// Currency object of a Price.
type PriceCurrency struct {
	// ID of the currency.
	ID *int64 `json:"id"`

	// Whether the currency is master or not, master meaning the default currency.
	Master *bool `json:"master"`

	// 3 characters identifier.
	Identifier string `json:"identifier"`

	// Name of the currency.
	Name *string `json:"name"`

	// How many decimals the currency has.
	Decimals int32 `json:"decimals"`
}

// An additional pricing.
type Additional struct {
	// Type of the additional pricing. absolute means the value is a price.
	Type AdditionalType `json:"type"`
	// The value of this additional pricing. Both absolute and percentage values here are represented as a money object.
	Value Price `json:"value"`
	// A title for this additional pricing.
	Title *string `json:"title"`
}

// Type of an additional pricing.
type AdditionalType string

// Know [AdditionalType] values.
const (
	AdditionalTypeAbsolute   AdditionalType = "absolute"
	AdditionalTypePercentage AdditionalType = "percentage"
)

// Objects a price group is applied to.
type AppliedTo struct {
	// ID's of the ChargePoints this price group has been applied to.
	ChargePoints []int64 `json:"chargePoints"`

	// ID's of the Sites this price group has been applied to.
	Sites []int64 `json:"sites"`

	// ID's of the Team Members this price group has been applied to.
	TeamMembers []int64 `json:"teamMembers"`
}
