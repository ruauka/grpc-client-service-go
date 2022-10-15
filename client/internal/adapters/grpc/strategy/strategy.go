// Package strategy - Package strategy
package strategy

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"client/internal/config"
	"client/internal/entities"

	"client/pb"

	"client/internal/utils/functions"
)

// Strategy - Interface for strategy struct.
type Strategy interface {
	HealthCheck(ctx context.Context) error
	Execute(ctx context.Context, req *entities.Request) (*entities.Response, error)
}

// strategy - structure for connecting to the gRPC service.
type strategy struct {
	// conn   *grpc.ClientConn
	client pb.StrategyClient
}

// NewStrategy - builder for connecting to the gRPC service.
func NewStrategy(config *config.Config) (Strategy, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", config.Host, config.Port),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("Connect to Strategy: failure")
		log.Println(err)
		return nil, err
	}
	log.Println("Connect to Strategy: ok")

	return &strategy{
		// conn:   conn,
		client: pb.NewStrategyClient(conn),
	}, nil
}

// NewMockStrategy - builder for mocking connecting to the gRPC service.
func NewMockStrategy() Strategy {
	testServer := &functions.TestGrpcServer{}
	testServer.StartStrategyServer(NewMock())

	return &strategy{
		//conn: testServer.Conn,
		client: testServer.ConnectStrategy(),
	}
}

// HealthCheck - check grpc service health.
func (s strategy) HealthCheck(ctx context.Context) error {
	_, err := s.client.HealthCheck(ctx, &emptypb.Empty{})
	if err != nil {
		errStatus, _ := status.FromError(err)
		log.Println(errStatus.Message())
		log.Println(errStatus.Code())
		return err
	}

	return nil
}

// Execute - main PRC.
func (s strategy) Execute(ctx context.Context, req *entities.Request) (*entities.Response, error) {
	payload := functions.FmtToGRPC(req)

	resp, err := s.client.Execute(ctx, payload)
	if err != nil {
		errStatus, _ := status.FromError(err)
		log.Println(errStatus.Message())
		log.Println(errStatus.Code())
		return nil, err
	}

	return functions.FmtFromGRPC(resp), nil
}
