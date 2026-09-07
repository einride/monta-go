package monta

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestListCharges_OperatorRole(t *testing.T) {
	owner := OperatorRoleOwner
	payer := OperatorRolePayer

	testCases := []struct {
		name         string
		operatorRole *OperatorRole
		wantParam    string
	}{
		{name: "omitted", operatorRole: nil},
		{name: "owner", operatorRole: &owner, wantParam: "owner"},
		{name: "payer", operatorRole: &payer, wantParam: "payer"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var gotQuery url.Values
			client := NewClient(WithToken(&Token{
				AccessToken:                "test-token",
				AccessTokenExpirationTime:  time.Now().Add(time.Hour),
				RefreshTokenExpirationTime: time.Now().Add(time.Hour),
			})).(*clientImpl)
			client.config.maxRetries = 0
			client.httpClient = &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "/api/v1/charges", r.URL.Path)
				gotQuery = r.URL.Query()
				return newTestResponse(http.StatusOK, `{"data":[],"meta":{}}`), nil
			})}

			_, err := client.ListCharges(context.Background(), &ListChargesRequest{
				OperatorRole: testCase.operatorRole,
			})
			assert.NilError(t, err)

			if testCase.wantParam == "" {
				assert.Assert(t, !gotQuery.Has("operatorRole"))
			} else {
				assert.Equal(t, testCase.wantParam, gotQuery.Get("operatorRole"))
			}
		})
	}
}
