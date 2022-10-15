package response

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"service/internal/logic"
	"service/internal/request"
	"service/internal/utils/dictionary"
)

func TestNewResponse(t *testing.T) {
	data := NewTestData()
	expected := &Response{
		Version:     "v.1.0.0",
		ExecuteDate: time.Now().Format(dictionary.LayoutToString),
		Results: []Result{
			{
				data.Result,
			},
		},
	}
	actual := NewResponse(data)
	require.Equal(t, expected, actual)
}

func NewTestData() *logic.Data {
	return &logic.Data{Request: *NewTestReq()}
}

func NewTestReq() *request.Request {
	var paymentDateBalance int32 = 0

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
				Payment:           0,
				PaymentDate:       time.Time{},
				ActualPaymentDate: time.Time{},
			},
		},
	}
}
