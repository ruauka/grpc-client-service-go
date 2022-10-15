package strategy

import (
	"context"
	"fmt"

	"client/pb"
)

// strategy - структура для подключения к тестовому gRPC-сервису.
type strategyMock struct {
	pb.UnimplementedStrategyServer
}

// NewMock - конструктор.
func NewMock() pb.StrategyServer {
	return &strategyMock{}
}

// Execute - мок заглушка для вызова gRPC процедуры и хендлера Call.
func (s *strategyMock) Execute(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	if req.Name == "err" {
		return nil, fmt.Errorf("execute call err")
	}

	return &pb.Response{
		Version:     "test Version",
		ExecuteDate: "test ExecuteDate",
	}, nil
}
