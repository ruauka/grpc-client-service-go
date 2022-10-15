package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"

	s "client/internal/adapters/grpc/strategy"
)

func TestHandler_Call(t *testing.T) {
	TestTable := []struct {
		jsonIn   []byte
		expected int
		testName string
	}{
		{
			jsonIn: []byte(`{
				"name": "Ivan",
				"surname": "Ivanov",
				"account": [
					{
						"payment_date_balance": 0
					}
				],
				"loan":[
					{
						"payment": 50,
						"payment_date": "2021-01-15T00:00:00Z",
						"actual_payment_date": "2021-01-16T00:00:00Z"
					}
				]
			}`),
			expected: http.StatusOK,
			testName: "test-1-good",
		},
		{
			jsonIn: []byte(`{
				"name": "Ivan",
				}`),
			expected: http.StatusBadRequest,
			testName: "test-2-bad-decode-400",
		},
		{
			jsonIn: []byte(`{
				"name": "Ivan"
				}`),
			expected: http.StatusBadRequest,
			testName: "test-2-bad-validate-400",
		},
		{
			jsonIn: []byte(`{
				"name": "err",
				"surname": "Ivanov",
				"account": [
					{
						"payment_date_balance": 0
					}
				],
				"loan":[
					{
						"payment": 50,
						"payment_date": "2021-01-15T00:00:00Z",
						"actual_payment_date": "2021-01-16T00:00:00Z"
					}
				]
			}`),
			expected: http.StatusBadRequest,
			testName: "test-4-gRPC-err-resp-400",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			payload := bytes.NewBuffer(testCase.jsonIn)
			r, _ := http.NewRequest("POST", "/execute", payload)
			w := httptest.NewRecorder()
			strategy := s.NewMockStrategy()
			handler := NewHandler(strategy)
			handler.Call(w, r, httprouter.Params{})
			require.Equal(t, testCase.expected, w.Code)
		})
	}
}
