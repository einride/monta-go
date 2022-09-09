package monta

// Location represents a geographical location.
type Location struct {
	// Coordinates of the location.
	Coordinates *LatLng `json:"coordinates"`
	// Address of the location.
	Address *Address `json:"address"`
}

// LatLng is a latitude longitude pair of geographical coordinates.
type LatLng struct {
	// Latitude of the coordinate.
	Latitude float64 `json:"latitude"`
	// Longitude of the coordinate.
	Longitude float64 `json:"longitude"`
}

// Address represents a postal address.
type Address struct {
	// Address1 is the first line of address.
	Address1 string `json:"address1"`

	// Address2 is the second line of address (optional).
	Address2 string `json:"address2"`

	// Address3 is the third line of address (optional).
	Address3 string `json:"address3"`

	// Zip is the zip code of the address.
	Zip string `json:"zip"`

	// City is the human-readable name of the city.
	City string `json:"city"`

	// Country is the human-readable name of the country.
	Country string `json:"country"`
}
