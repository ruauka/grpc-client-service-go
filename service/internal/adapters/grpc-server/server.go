// Package grpc_server - Package grpc_server
package grpc_server //nolint:revive,stylecheck

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"service/internal/strategy/logic"
	"service/internal/strategy/response"
	"service/pb"
)

// Server - server struct.
type Server struct {
	pb.UnimplementedStrategyServer
	GrpcServer *grpc.Server
}

// NewGRPCServer - server builder.
func NewGRPCServer() *Server {
	return &Server{
		GrpcServer: grpc.NewServer(),
	}
}

// Execute - main procedure.
func (s *Server) Execute(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("new request")
	// covert from client to service format
	message := FmtFromGRPC(req)
	// create data dict
	data := logic.NewData(&message)
	// script logic call
	data.LocalCount()
	data.ResultCount()
	// create response
	respGrpc := response.NewResponse(data)
	// covert from service to client format
	respClient := FmtToGRPC(&respGrpc)
	return &respClient, nil
}

// HealthCheck - gRPC connect alive check.
func (s *Server) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	log.Println("ping health check")

	return &emptypb.Empty{}, nil
}
