package functions

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/credentials/insecure"

	"client/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufnet = "bufnet"

const bufSize = 1024 * 1024

// TestGrpcServer implements test grpc server with underlying bufconn.
// initial structure is ready to use.
type TestGrpcServer struct {
	lis    *bufconn.Listener
	server *grpc.Server
	Conn   *grpc.ClientConn
}

// StartStrategyServer creates and starts grpc server with given API for testing purposes.
func (t *TestGrpcServer) StartStrategyServer(api pb.StrategyServer) {
	// create listener with bufconn (no OS sockets)
	t.lis = bufconn.Listen(bufSize)

	// create new server
	t.server = grpc.NewServer()

	// register service with this server
	pb.RegisterStrategyServer(t.server, api)

	// start server in separate goroutine
	go func() {
		if err := t.server.Serve(t.lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

// ConnectStrategy connects to a test server and returns client API instance.
func (t *TestGrpcServer) ConnectStrategy() pb.StrategyClient {
	var err error

	// connect to server
	log.Printf("connecting to %s", bufnet)
	t.Conn, err = grpc.DialContext(context.Background(),
		bufnet,
		grpc.WithContextDialer(t.bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Failed to dial %s: %v", bufnet, err)
	}

	// create client
	return pb.NewStrategyClient(t.Conn)
}

func (t *TestGrpcServer) bufDialer(context.Context, string) (net.Conn, error) {
	return t.lis.Dial()
}

// StopServer .
func (t *TestGrpcServer) StopServer() {
	t.server.Stop()
}

// CloseConnection closes underlying server connection.
func (t *TestGrpcServer) CloseConnection() {
	_ = t.Conn.Close() //nolint:errcheck
}
