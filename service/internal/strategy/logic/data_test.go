package logic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"service/internal/strategy/request"
)

func TestNewData(t *testing.T) {
	var paymentDateBalance int32 = 0

	expected := &Data{
		Request: request.Request{
			Name:    "Ivan",
			Surname: "Ivanov",
			Accounts: []request.Account{
				{
					PaymentDateBalance: &paymentDateBalance,
				},
			},
			Loans: []request.Loan{
				{
					Payment:           0,
					PaymentDate:       time.Time{},
					ActualPaymentDate: time.Time{},
				},
			},
		},
		Delinquency: Delinquency{},
		Result:      Result{},
	}
	actual := NewData(NewTestReq())
	require.Equal(t, expected, actual)
}
