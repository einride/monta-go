package monta

import (
	"net/url"
	"strconv"
)

// Common page filters.
type PageFilters struct {
	// Page number to request (starts with 0).
	Page int
	// Number of items PerPage (between 1 and 100, default 10).
	PerPage int
}

// Set the filters to the given query.
func (p PageFilters) Apply(query url.Values) {
	if p.Page > 0 {
		query.Set("page", strconv.Itoa(p.Page))
	}
	if p.PerPage > 0 {
		query.Set("perPage", strconv.Itoa(p.PerPage))
	}
}
