package monta

import (
	"encoding/json"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestWebhookEvent_MarshalJSON(t *testing.T) {
	expected := strings.TrimSpace(`
{
  "id": 1,
  "consumerId": 1,
  "operatorId": 1,
  "eventType": "charges",
  "payload": {
    "entityType": "charge",
    "entityId": "1",
    "eventType": "updated",
    "payload": {}
  },
  "status": "completed",
  "error": "",
  "createdAt": "2022-02-04T14:50:33Z",
  "updatedAt": "2022-02-04T14:50:36Z"
}
`)
	var webhookEvent WebhookEvent
	assert.NilError(t, json.Unmarshal([]byte(expected), &webhookEvent))
	actual, err := json.MarshalIndent(&webhookEvent, "", "  ")
	assert.NilError(t, err)
	assert.Equal(t, expected, string(actual))
}
