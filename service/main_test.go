package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"service/pb"

	"service/internal/adapters/grpc-server"
)

func BenchmarkExecute(b *testing.B) {
	srv := grpc_server.NewGRPCServer()

	ctx := context.Background()
	payload := prepare()

	for i := 0; i < b.N; i++ {
		srv.Execute(ctx, payload)
	}
}

func prepare() *pb.Request {
	data, err := ioutil.ReadFile("testdata/input.json")
	if err != nil {
		fmt.Println(err)
	}

	var payload pb.Request
	json.Unmarshal(data, &payload)

	return &payload
}
