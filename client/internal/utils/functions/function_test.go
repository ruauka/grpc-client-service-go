package functions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"client/internal/request"
	"client/internal/response"
	"client/pb"
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
	req := NewHttpReq()

	t.Run("validation error", func(t *testing.T) {
		err := Validate(req)
		require.EqualError(t, err, "Key: 'Request.Loans[0].Payment' Error:Field validation for 'Payment' failed on the 'required' tag")
	})
	t.Run("success validation", func(t *testing.T) {
		req.Loans[0].Payment = 10
		err := Validate(req)
		require.Equal(t, nil, err)
	})
}

func TestFmtToGRPC(t *testing.T) {
	req := NewHttpReq()
	expected := NewGRPCRequest()

	t.Run("test 1", func(t *testing.T) {
		actual := FmtToGRPC(req)
		require.Equal(t, expected, actual)
	})
}

func TestFmtFromGRPC(t *testing.T) {
	resp := NewGRPCResponse()
	expected := NewHttpResponse()

	t.Run("test 1", func(t *testing.T) {
		actual := FmtFromGRPC(resp)
		require.Equal(t, expected, actual)
	})
}

func NewHttpReq() *request.Request {
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

func NewGRPCRequest() *pb.Request {
	return &pb.Request{
		Name:    "Ivan",
		Surname: "Ivanov",
		Account: []*pb.Account{
			{
				PaymentDateBalance: 10,
			},
		},
		Loan: []*pb.Loan{
			{
				PaymentDate:       timestamppb.New(time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC)),
				ActualPaymentDate: timestamppb.New(time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC)),
			},
		},
	}
}

func NewHttpResponse() *response.Response {
	return &response.Response{
		Version:     "1.0",
		ExecuteDate: "11-11-11",
		Results: []response.Result{
			{
				EnoughMoneyByMonths:          [6]int32{1, 0, 0, 0, 1, 1},
				DelinquencyByMonths:          [6]int32{1, 0, 0, 0, 1, 1},
				DelinquencyDurationDaysTotal: 10,
				DelinquencySumTotal:          200,
			},
		},
	}
}

func NewGRPCResponse() *pb.Response {
	return &pb.Response{
		Version:     "1.0",
		ExecuteDate: "11-11-11",
		Results: []*pb.Result{
			{
				EnoughMoneyByMonths:          []int32{1, 0, 0, 0, 1, 1},
				DelinquencyByMonths:          []int32{1, 0, 0, 0, 1, 1},
				DelinquencyDurationDaysTotal: 10,
				DelinquencySumTotal:          200,
			},
		},
	}
}
