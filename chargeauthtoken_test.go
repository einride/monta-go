package monta

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestChargeAuthToken_MarshalJSON(t *testing.T) {
	expected := strings.TrimSpace(`
{
  "id": 1,
  "identifier": "38C58DB85F4",
  "type": "vehicleId",
  "teamId": 13,
  "name": "Monta Team Key",
  "partnerExternalId": "abc",
  "blockedAt": "2022-05-12T15:56:45.99Z",
  "createdAt": "2022-05-12T15:56:45.99Z",
  "updatedAt": "2022-05-12T16:56:45.99Z"
}
	`)
	var token ChargeAuthToken
	assert.NilError(t, json.Unmarshal([]byte(expected), &token))
	actual, err := json.MarshalIndent(&token, "", "  ")
	assert.NilError(t, err)
	assert.Equal(t, expected, string(actual))
}

func TestChargeAuthToken_JSONEncode(t *testing.T) {
	request := CreateChargeAuthTokenRequest{
		TeamID:         239264,
		UserID:         nil,
		Identifier:     "testkey",
		Type:           "rfid",
		Name:           new(string),
		MontaNetwork:   false,
		RoamingNetwork: false,
	}
	var requestBody bytes.Buffer
	err := json.NewEncoder(&requestBody).Encode(&request)
	assert.NilError(t, err)
	var tokenRequest CreateChargeAuthTokenRequest
	res, err := doPatch[ChargeAuthToken](context.TODO(), c, path, &requestBody)
	c.Clie

	// assert.NilError(t, json.Unmarshal(requestBody, &tokenRequest))
	// actual, err := json.MarshalIndent(&token, "", "  ")
	// assert.NilError(t, err)
	// assert.Equal(t, expected, string(actual))
}
