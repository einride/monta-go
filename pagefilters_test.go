package monta_test

import (
	"net/url"
	"testing"

	"go.einride.tech/monta"
	"gotest.tools/v3/assert"
)

func Test_PageFilters_Apply_Empty(t *testing.T) {
	query := url.Values{}
	monta.PageFilters{}.Apply(query)
	assert.DeepEqual(t, query, url.Values{})
}

func Test_PageFilters_Apply_NotEmpty(t *testing.T) {
	query := url.Values{}

	monta.PageFilters{
		Page:    1,
		PerPage: 2,
	}.Apply(query)

	expected := url.Values{}
	expected.Set("page", "1")
	expected.Set("perPage", "2")
	assert.DeepEqual(t, query, expected)
}
