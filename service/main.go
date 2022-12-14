// Package main
package main

//nolint:goimports
import (
	"log"
	"net"

	"service/internal/adapters/grpc-server"
	"service/pb"
)

func main() {
	// crete gRPC server
	srv := grpc_server.NewGRPCServer()
	// register gRPC server
	pb.RegisterStrategyServer(srv.GrpcServer, srv)
	// listen net
	l, err := net.Listen("tcp", ":8080") //nolint:gosec
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server...")

	// start gRPC server
	if err := srv.GrpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
