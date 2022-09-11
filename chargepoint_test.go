package monta

import (
	"encoding/json"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestChargePoint_MarshalJSON(t *testing.T) {
	expected := strings.TrimSpace(`
{
  "id": 1,
  "siteId": 42,
  "serialNumber": "string",
  "name": "Monta CPH HQ",
  "visibility": "public",
  "maxKW": 150,
  "type": "ac",
  "note": "In order to access this charge point enter 0000 as code at the gate.",
  "state": "available",
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
  "deeplinks": {
    "app": "http://example.com",
    "web": "http://example.com"
  }
}
	`)
	var chargePoint ChargePoint
	assert.NilError(t, json.Unmarshal([]byte(expected), &chargePoint))
	actual, err := json.MarshalIndent(&chargePoint, "", "  ")
	assert.NilError(t, err)
	assert.Equal(t, expected, string(actual))
}
