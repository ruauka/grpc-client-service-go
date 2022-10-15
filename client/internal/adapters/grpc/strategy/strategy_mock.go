package strategy

import (
	"context"
	"fmt"

	"client/pb"
)

// strategy - structure for connecting to the test gRPC service. For Mock.
type strategyMock struct {
	pb.UnimplementedStrategyServer
}

// NewMock - Builder for strategyMock.
func NewMock() pb.StrategyServer {
	return &strategyMock{}
}

// Execute - mock plug for main RPC and handler 'Call'.
func (s *strategyMock) Execute(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	if req.Name == "err" {
		return nil, fmt.Errorf("execute call err")
	}

	return &pb.Response{
		Version:     "test Version",
		ExecuteDate: "test ExecuteDate",
	}, nil
}
