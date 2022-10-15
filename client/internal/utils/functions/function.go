// Package functions - a package of help functions.
package functions

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/timestamppb"

	"client/internal/entities"
	"client/internal/utils/dictionary"
	"client/pb"
)

// MakeJSONResponse make http response to client.
func MakeJSONResponse(w http.ResponseWriter, httpStatus int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(resp) //nolint:errcheck,gosec
}

// Validate - Request struct validate.
func Validate(req *entities.Request) error {
	err := dictionary.Validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}

	return nil
}

// DateDiff - calculation of the difference in dates..
func DateDiff(date1, date2 time.Time) int32 {
	duration := date2.Sub(date1).Hours() / dictionary.HourDivider
	return int32(duration)
}

// FmtToGRPC - convert from *request.Request to *pb.Request.
func FmtToGRPC(req *entities.Request) *pb.Request {
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

// FmtFromGRPC - convert from *pb.Response to *response.Response.
func FmtFromGRPC(resp *pb.Response) *entities.Response {
	var results []entities.Result
	enoughMoneyByMonths := [6]int32{}
	delinquencyByMonths := [6]int32{}

	for _, v := range resp.Results {
		for i, j := range v.EnoughMoneyByMonths { //nolint:gosimple
			enoughMoneyByMonths[i] = j
		}
		for k, n := range v.DelinquencyByMonths { //nolint:gosimple
			delinquencyByMonths[k] = n
		}

		result := &entities.Result{
			EnoughMoneyByMonths:          enoughMoneyByMonths,
			DelinquencyByMonths:          delinquencyByMonths,
			DelinquencyDurationDaysTotal: v.DelinquencyDurationDaysTotal,
			DelinquencySumTotal:          v.DelinquencySumTotal,
		}
		results = append(results, *result)
	}

	return &entities.Response{
		Version:     resp.Version,
		ExecuteDate: resp.ExecuteDate,
		Results:     results,
	}
}
