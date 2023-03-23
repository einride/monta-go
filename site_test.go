package monta

import (
	"encoding/json"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestSite_MarshalJSON(t *testing.T) {
	expected := strings.TrimSpace(`
{
  "id": 1,
  "name": "Monta CPH HQ",
  "chargePointCount": 42,
  "activeChargePointCount": 33,
  "availableChargePointCount": 4,
  "maxKW": 150,
  "type": "ac",
  "visibility": "public",
  "note": "In order to access this site enter 0000 as code at the gate.",
  "partnerExternalId": "abc",
  "location": {
    "coordinates": {
      "latitude": 55.6760968,
      "longitude": 12.5683371
    },
    "address": {
      "address1": "Strandboulevarden 122",
      "address2": "København",
      "address3": "5. sal",
      "zip": "2100",
      "city": "Copenhagen",
      "country": "Denmark"
    }
  },
  "connectors": [
    {
      "identifier": "ccs",
      "name": "CCS"
    }
  ],
  "createdAt": "2022-05-12T15:56:45.999189Z",
  "updatedAt": "2022-05-17T15:56:45.999189Z"
}
	`)
	var site Site
	assert.NilError(t, json.Unmarshal([]byte(expected), &site))
	actual, err := json.MarshalIndent(&site, "", "  ")
	assert.NilError(t, err)
	assert.Equal(t, expected, string(actual))
}
