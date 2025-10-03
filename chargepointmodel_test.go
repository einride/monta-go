package monta

import (
	"encoding/json"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestChargePointModel_MarshalJSON(t *testing.T) {
	expected := strings.TrimSpace(`
{
  "id": 1,
  "identifier": "rolec_securicharge_ev_dual",
  "name": "SecuriCharge Dual",
  "brand": {
    "id": 1,
    "name": "Rolec"
  },
  "features": [
    {
      "key": "firmware_management",
      "description": "Firmware updates remotely",
      "requirements": "Support UpdateFirmware functionality",
      "enabled": true
    }
  ],
  "createdAt": "2022-05-12T15:56:45.99Z",
  "updatedAt": "2022-05-12T16:56:45.99Z"
}`)
	var chargePointModel ChargePointModel
	assert.NilError(t, json.Unmarshal([]byte(expected), &chargePointModel))
	actual, err := json.MarshalIndent(&chargePointModel, "", "  ")
	assert.NilError(t, err)
	assert.Equal(t, expected, string(actual))
}
