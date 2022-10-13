package main

import (
	"log"
	"net"
	_ "net/http/pprof"

	"google.golang.org/grpc"

	"service/grpc_server"
	"service/pb"
)

func main() {
	s := grpc.NewServer()
	srv := &grpc_server.Server{}
	pb.RegisterExecutorServer(s, srv)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
