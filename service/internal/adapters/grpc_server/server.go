package grpc_server

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"service/internal/strategy/logic"
	"service/internal/strategy/response"
	"service/pb"
)

type Server struct {
	GrpcServer *grpc.Server
	pb.UnimplementedStrategyServer
}

func NewGRPCServer() *Server {
	return &Server{
		GrpcServer: grpc.NewServer(),
	}
}

func (s *Server) Execute(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("get new request")

	message := FmtFromGRPC(req)
	// создание словаря данных.
	data := logic.NewData(message)
	// вызов методов логики.
	data.LocalCount()
	data.ResultCount()
	// создание объекта ответа.
	resp := response.NewResponse(data)

	return FmtToGRPC(resp), nil
}

func (s *Server) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	log.Println("ping health check")

	return &emptypb.Empty{}, nil
}
