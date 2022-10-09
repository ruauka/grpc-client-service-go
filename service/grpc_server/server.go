package grpc_server

import (
	"context"
	"fmt"

	"service/internal/logic"
	"service/internal/request"
	"service/internal/response"
	"service/pb"
)

type Server struct {
	pb.UnimplementedExecutorServer
}

func (s *Server) Execute(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	fmt.Println("ping")
	message := FmtFromGRPC(req)
	// создание словаря данных.
	data := logic.NewData(message)
	// вызов методов логики.
	data.LocalCount()
	data.ResultCount()
	// создание объекта ответа.
	resp := response.NewResponse(data)

	fmt.Println(resp)

	return FmtToGRPC(resp), nil
}

func FmtToGRPC(resp *response.Response) *pb.Response {
	var results []*pb.Result
	enoughMoneyByMonths := make([]int32, len(resp.Results[0].EnoughMoneyByMonths))
	delinquencyByMonths := make([]int32, len(resp.Results[0].DelinquencyByMonths))

	for _, v := range resp.Results {
		for i, j := range v.EnoughMoneyByMonths {
			enoughMoneyByMonths[i] = j
		}
		for k, n := range v.DelinquencyByMonths {
			delinquencyByMonths[k] = n
		}

		result := &pb.Result{
			EnoughMoneyByMonths:          enoughMoneyByMonths,
			DelinquencyByMonths:          delinquencyByMonths,
			DelinquencyDurationDaysTotal: v.DelinquencyDurationDaysTotal,
			DelinquencySumTotal:          v.DelinquencySumTotal,
		}
		results = append(results, result)
	}

	return &pb.Response{
		Version:     resp.Version,
		ExecuteDate: resp.ExecuteDate,
		Results:     results,
	}
}

func FmtFromGRPC(req *pb.Request) *request.Request {
	accounts := make([]request.Account, len(req.Accounts))
	for i, v := range req.Accounts {
		accounts[i].PaymentDateBalance = &v.PaymentDateBalance
	}

	loans := make([]request.Loan, len(req.Loans))
	for i, v := range req.Loans {
		loans[i].Payment = v.Payment
		loans[i].PaymentDate = v.PaymentDate.AsTime()
		loans[i].ActualPaymentDate = v.ActualPaymentDate.AsTime()
	}

	return &request.Request{
		Name:     req.Name,
		Surname:  req.Surname,
		Accounts: accounts,
		Loans:    loans,
	}
}

//
//i, _ := json.MarshalIndent(request, "", "   ")
//fmt.Println(string(i))
