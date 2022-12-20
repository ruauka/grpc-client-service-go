package grpc_server //nolint:revive,stylecheck

import (
	"service/internal/strategy/request"
	"service/internal/strategy/response"
	"service/pb"
)

// FmtToGRPC - convert from *response.Response to pb.Response.
func FmtToGRPC(resp *response.Response) pb.Response {
	results := make([]*pb.Result, len(resp.Results))
	enoughMoneyByMonths := make([]int32, len(resp.Results[0].EnoughMoneyByMonths))
	delinquencyByMonths := make([]int32, len(resp.Results[0].DelinquencyByMonths))

	for index, v := range resp.Results {
		for i, j := range v.EnoughMoneyByMonths { //nolint:gosimple
			enoughMoneyByMonths[i] = j
		}
		for k, n := range v.DelinquencyByMonths { //nolint:gosimple
			delinquencyByMonths[k] = n
		}

		result := &pb.Result{
			EnoughMoneyByMonths:          enoughMoneyByMonths,
			DelinquencyByMonths:          delinquencyByMonths,
			DelinquencyDurationDaysTotal: v.DelinquencyDurationDaysTotal,
			DelinquencySumTotal:          v.DelinquencySumTotal,
		}

		results[index] = result
	}

	return pb.Response{
		Version:     resp.Version,
		ExecuteDate: resp.ExecuteDate,
		Results:     results,
	}
}

// FmtFromGRPC - convert from *pb.Request to request.Request.
func FmtFromGRPC(req *pb.Request) request.Request {
	accounts := make([]request.Account, len(req.Account))
	for i, v := range req.Account {
		accounts[i].PaymentDateBalance = &v.PaymentDateBalance
	}

	loans := make([]request.Loan, len(req.Loan))
	for i, v := range req.Loan {
		loans[i].Payment = v.Payment
		loans[i].PaymentDate = v.PaymentDate.AsTime()
		loans[i].ActualPaymentDate = v.ActualPaymentDate.AsTime()
	}

	return request.Request{
		Name:     req.Name,
		Surname:  req.Surname,
		Accounts: accounts,
		Loans:    loans,
	}
}
