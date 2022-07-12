package integration

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/thalissonfelipe/banking/banking/domain/vos"
	"github.com/thalissonfelipe/banking/banking/tests/dockertest"
	"github.com/thalissonfelipe/banking/banking/tests/fakes"
	"github.com/thalissonfelipe/banking/banking/tests/testdata"
	"github.com/thalissonfelipe/banking/banking/tests/testenv"
)

func TestIntegration_GetAccountBalance(t *testing.T) {
	testCases := []struct {
		name           string
		uriSetup       func(t *testing.T) string
		expectedStatus int
	}{
		{
			name: "should return status code 200 if account exists",
			uriSetup: func(t *testing.T) string {
				acc := createAccount(t, testdata.GetValidCPF(), testdata.GetValidSecret(), 0)

				return fmt.Sprintf("%s/api/v1/accounts/%s/balance", testenv.ServerURL, acc.ID.String())
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "should return status code 404 if account does not exist",
			uriSetup: func(t *testing.T) string {
				return fmt.Sprintf("%s/api/v1/accounts/%s/balance", testenv.ServerURL, vos.NewAccountID().String())
			},
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			request := fakes.FakeRequest(http.MethodGet, tt.uriSetup(t), nil)
			resp, err := http.DefaultClient.Do(request)
			require.NoError(t, err)

			defer resp.Body.Close()

			var body bytes.Buffer

			_, err = io.Copy(&body, resp.Body)
			require.NoError(t, err)

			t.Log(body.String())

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			dockertest.TruncateTables(context.Background(), testenv.DB)
		})
	}
}