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
  "teamId": 1,
  "serialNumber": "string",
  "name": "Monta CPH HQ",
  "visibility": "public",
  "maxKW": 150,
  "modelId": 0,
  "type": "ac",
  "note": "In order to access this charge point enter 0000 as code at the gate.",
  "state": "available",
  "lastMeterReadingKwh": 913.2,
  "cablePluggedIn": true,
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
  "deeplinks": {
    "app": "http://example.com",
    "web": "http://example.com"
  },
  "createdAt": "2022-05-12T15:56:45.999189Z",
  "updatedAt": "2022-05-17T15:56:45.999189Z",
  "operator": {
    "id": 445,
    "name": "Einride",
    "identifier": "einride",
    "vatNumber": "123",
    "partnerId": 423
  }
}
	`)
	var chargePoint ChargePoint
	assert.NilError(t, json.Unmarshal([]byte(expected), &chargePoint))
	actual, err := json.MarshalIndent(&chargePoint, "", "  ")
	assert.NilError(t, err)
	assert.Equal(t, expected, string(actual))
}
