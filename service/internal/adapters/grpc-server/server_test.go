package grpc_server

import (
	"context"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"service/internal/strategy/utils/dictionary"
	"service/pb"

	"github.com/stretchr/testify/require"
)

func TestServer_Execute(t *testing.T) {
	srv := NewGRPCServer()
	req := &pb.Request{}
	expected := &pb.Response{
		Version:     "v.1.0.0",
		ExecuteDate: time.Now().Format(dictionary.LayoutToString),
		Results: []*pb.Result{
			{
				EnoughMoneyByMonths: []int32{0, 0, 0, 0, 0, 0},
				DelinquencyByMonths: []int32{0, 0, 0, 0, 0, 0},
			},
		},
	}

	actual, _ := srv.Execute(context.Background(), req)

	require.Equal(t, expected, actual)
}

func TestServer_HealthCheck(t *testing.T) {
	srv := NewGRPCServer()
	expected := &emptypb.Empty{}

	actual, _ := srv.HealthCheck(context.Background(), nil)

	require.Equal(t, expected, actual)
}
