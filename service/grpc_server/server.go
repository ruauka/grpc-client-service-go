package grpc_server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

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

func (s *Server) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
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
