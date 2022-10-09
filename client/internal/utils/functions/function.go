// Package functions - пакет вспомогательных функций.
package functions

import (
	"time"

	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/timestamppb"

	"client/internal/request"
	"client/internal/response"
	"client/internal/utils/dictionary"
	"client/pb"
)

// Validate - валидация структуры Request.
func Validate(req *request.Request) error {
	err := dictionary.Validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}

	return nil
}

// DateDiff - вспомогательная функция расчета разницы в датах.
func DateDiff(date1, date2 time.Time) int32 {
	duration := date2.Sub(date1).Hours() / dictionary.HourDivider
	return int32(duration)
}

func FmtFromGRPC(resp *pb.Response) *response.Response {
	var results []response.Result
	enoughMoneyByMonths := [6]int32{}
	delinquencyByMonths := [6]int32{}

	for _, v := range resp.Results {
		for i, j := range v.EnoughMoneyByMonths {
			enoughMoneyByMonths[i] = j
		}
		for k, n := range v.DelinquencyByMonths {
			delinquencyByMonths[k] = n
		}

		result := &response.Result{
			EnoughMoneyByMonths:          enoughMoneyByMonths,
			DelinquencyByMonths:          delinquencyByMonths,
			DelinquencyDurationDaysTotal: v.DelinquencyDurationDaysTotal,
			DelinquencySumTotal:          v.DelinquencySumTotal,
		}
		results = append(results, *result)
	}

	return &response.Response{
		Version:     resp.Version,
		ExecuteDate: resp.ExecuteDate,
		Results:     results,
	}
}

func FmtToGRPC(req *request.Request) *pb.Request {
	var accountsSl []*pb.Account
	for _, v := range req.Accounts {
		account := &pb.Account{PaymentDateBalance: *v.PaymentDateBalance}
		accountsSl = append(accountsSl, account)
	}

	var loansSl []*pb.Loan
	for _, v := range req.Loans {
		loan := &pb.Loan{
			Payment:           v.Payment,
			PaymentDate:       timestamppb.New(v.PaymentDate),
			ActualPaymentDate: timestamppb.New(v.ActualPaymentDate),
		}
		loansSl = append(loansSl, loan)
	}

	return &pb.Request{
		Name:    req.Name,
		Surname: req.Surname,
		Account: accountsSl,
		Loan:    loansSl,
	}
}
