package monta

import (
	"encoding/json"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestWalletTransaction_MarshalJSON(t *testing.T) {
	expected := strings.TrimSpace(`
{
  "id": 1,
  "fromAmount": 13.77,
  "fromCurrency": {
    "identifier": "DKK",
    "name": "Danish krone",
    "decimals": 2
  },
  "fromType": "team",
  "from": {
    "id": 14,
    "publicName": "Monta",
    "partnerExternalId": "abcd"
  },
  "toAmount": 13.77,
  "toCurrency": {
    "identifier": "DKK",
    "name": "Danish krone",
    "decimals": 2
  },
  "toType": "operator",
  "to": {
    "id": 14,
    "name": "Monta",
    "identifier": "monta",
    "vatNumber": "FOO-123-ABC"
  },
  "exchangeRate": 1,
  "createdAt": "2022-04-22T09:47:05Z",
  "updatedAt": "2022-04-22T09:47:06Z",
  "referenceType": "SubscriptionPurchase",
  "group": "withdraw",
  "state": "complete",
  "note": "Test transaction."
}
	`)
	var walletTransaction WalletTransaction
	assert.NilError(t, json.Unmarshal([]byte(expected), &walletTransaction))
	expectedFromTeam := &Team{
		ID:                14,
		PublicName:        "Monta",
		PartnerExternalID: toPointer("abcd"),
	}
	assert.DeepEqual(t, expectedFromTeam, walletTransaction.FromTeam)
	walletTransaction.From = nil
	expectedToOperator := &Operator{
		ID:         14,
		Name:       "Monta",
		Identifier: "monta",
		VATNumber:  "FOO-123-ABC",
	}
	assert.DeepEqual(t, expectedToOperator, walletTransaction.ToOperator)
	walletTransaction.To = nil
	actual, err := json.MarshalIndent(&walletTransaction, "", "  ")
	assert.NilError(t, err)
	assert.Equal(t, expected, string(actual))
}

func toPointer[T any](v T) *T {
	return &v
}
