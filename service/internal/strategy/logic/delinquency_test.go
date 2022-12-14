package logic

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"service/internal/strategy/request"
)

func TestNewLocal(t *testing.T) {
	expected := &Delinquency{}
	actual := NewLocal()
	require.Equal(t, expected, actual)
}

func TestLocal_EnoughMoneyCount(t *testing.T) {
	data := NewTestData()

	TestTable := []struct {
		paymentDateBalance int32
		payment            int32
		expected           bool
		testName           string
	}{
		{
			paymentDateBalance: 100,
			payment:            50,
			expected:           true,
			testName:           "test-1-EnoughMoney=true",
		},
		{
			paymentDateBalance: 40,
			payment:            50,
			expected:           false,
			testName:           "test-2-EnoughMoney=false",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			data.Request.Accounts[0].PaymentDateBalance = &testCase.paymentDateBalance
			data.Request.Loans[0].Payment = testCase.payment
			data.EnoughMoneyCount(data)
			actual := data.Delinquency.EnoughMoney[0]
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestLocal_DelinquencyCount(t *testing.T) {
	data := NewTestData()

	TestTable := []struct {
		actualPaymentDate time.Time
		paymentDate       time.Time
		expected          bool
		testName          string
	}{
		{
			actualPaymentDate: time.Date(2021, 01, 15, 0, 0, 0, 0, time.UTC),
			paymentDate:       time.Date(2021, 01, 15, 0, 0, 0, 0, time.UTC),
			expected:          false,
			testName:          "test-1-Delinquency=false",
		},
		{
			actualPaymentDate: time.Date(2021, 01, 16, 0, 0, 0, 0, time.UTC),
			paymentDate:       time.Date(2021, 01, 15, 0, 0, 0, 0, time.UTC),
			expected:          true,
			testName:          "test-2-Delinquency=true",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			data.Request.Loans[0].ActualPaymentDate = testCase.actualPaymentDate
			data.Request.Loans[0].PaymentDate = testCase.paymentDate
			data.Delinquency.DelinquencyCount(data)
			actual := data.Delinquency.Delinquency[0]
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestLocal_DelinquencyDurationCount(t *testing.T) {
	data := NewTestData()

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
			testName:          "test-1-DelinquencyDurationDays=0",
		},
		{
			actualPaymentDate: time.Date(2021, 01, 20, 0, 0, 0, 0, time.UTC),
			paymentDate:       time.Date(2021, 01, 15, 0, 0, 0, 0, time.UTC),
			expected:          5,
			testName:          "test-2-DelinquencyDurationDays=5",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			data.Request.Loans[0].ActualPaymentDate = testCase.actualPaymentDate
			data.Request.Loans[0].PaymentDate = testCase.paymentDate
			data.Delinquency.DelinquencyDurationCount(data)
			actual := data.Delinquency.DelinquencyDurationDays[0]
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestLocal_DelinquencySumCount(t *testing.T) {
	data := NewTestData()

	TestTable := []struct {
		delinquencyDurationDays int32
		payment                 int32
		paymentDateBalance      int32
		expected                int32
		testName                string
	}{
		{
			delinquencyDurationDays: 0,
			payment:                 100,
			paymentDateBalance:      150,
			expected:                0,
			testName:                "test-1-DelinquencySum=0",
		},
		{
			delinquencyDurationDays: 5,
			payment:                 100,
			paymentDateBalance:      90,
			expected:                10,
			testName:                "test-2-DelinquencySum=10",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			data.Delinquency.DelinquencyDurationDays[0] = testCase.delinquencyDurationDays
			data.Request.Loans[0].Payment = testCase.payment
			data.Request.Accounts[0].PaymentDateBalance = &testCase.paymentDateBalance
			data.DelinquencySumCount(data)
			actual := data.DelinquencySum[0]
			fmt.Println(data.Delinquency.DelinquencySum)
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func NewTestData() *Data {
	return &Data{Request: *NewTestReq()}
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
