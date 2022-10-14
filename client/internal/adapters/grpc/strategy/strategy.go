// Package strategy - Package strategy
package strategy

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"client/internal/entities"

	"client/pb"

	"client/internal/utils/functions"
)

// Strategy - Интерфейс для компонента Strategy.
type Strategy interface {
	HealthCheck(ctx context.Context) error
	Execute(ctx context.Context, req *entities.Request) (*entities.Response, error)
}

// strategy - структура для подключения к gRPC-сервису.
type strategy struct {
	// conn   *grpc.ClientConn
	client pb.ExecutorClient
}

// NewStrategy - конструктор подключения к gRPC-сервису.
func NewStrategy() (Strategy, error) {
	conn, err := grpc.Dial(
		":8080",
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
		//conn:   conn,
		client: pb.NewExecutorClient(conn),
	}, nil
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

// Execute - вызов основной rpc процедуры.
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
