package main

import (
	"log"
	"net"
	_ "net/http/pprof"

	"service/internal/adapters/grpc_server"
	"service/pb"
)

func main() {
	srv := grpc_server.NewGRPCServer()

	pb.RegisterStrategyServer(srv.GrpcServer, srv.UnimplementedStrategyServer)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server...")

	if err := srv.GrpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
