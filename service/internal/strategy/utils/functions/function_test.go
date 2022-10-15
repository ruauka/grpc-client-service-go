package functions

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"service/internal/strategy/request"
)

func TestDateDiff(t *testing.T) {
	TestTable := []struct {
		actualPaymentDate time.Time
		paymentDate       time.Time
		expected          int32
		testName          string
	}{
		{
			actualPaymentDate: time.Date(2021, 01, 15, 0, 0, 0, 0, time.UTC),
			paymentDate:       time.Date(2021, 01, 15, 0, 0, 0, 0, time.UTC),
			expected:          0,
			testName:          "test-1-diff=0",
		},
		{
			actualPaymentDate: time.Date(2021, 01, 20, 0, 0, 0, 0, time.UTC),
			paymentDate:       time.Date(2021, 01, 15, 0, 0, 0, 0, time.UTC),
			expected:          5,
			testName:          "test-2-diff=5",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := DateDiff(testCase.paymentDate, testCase.actualPaymentDate)
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestValidateStruct(t *testing.T) {
	req := NewTestReq()
	t.Run("validation error", func(t *testing.T) {
		err := Validate(req)
		fmt.Println(err)
		require.EqualError(t, err, "Key: 'Request.Loans[0].Payment' Error:Field validation for 'Payment' failed on the 'required' tag")
	})
	t.Run("success validation", func(t *testing.T) {
		req.Loans[0].Payment = 10
		err := Validate(req)
		require.Equal(t, nil, err)
	})
}

func NewTestReq() *request.Request {
	var paymentDateBalance int32 = 10

	return &request.Request{
		Name:    "Ivan",
		Surname: "Ivanov",
		Accounts: []request.Account{
			{
				PaymentDateBalance: &paymentDateBalance,
			},
		},
		Loans: []request.Loan{
			{

				PaymentDate:       time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
				ActualPaymentDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			},
		},
	}
}
