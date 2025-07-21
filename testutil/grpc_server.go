package testutil

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Simple gRPC server for testing
type testGRPCServer struct {
	UnimplementedGreeterServer
}

func (s *testGRPCServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return &HelloReply{
		Message: fmt.Sprintf("Hello %s", req.GetName()),
	}, nil
}

// StartGRPCServer starts a simple test gRPC server
func StartGRPCServer() (*grpc.Server, string, error) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, "", fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterGreeterServer(s, &testGRPCServer{})
	
	// Register reflection service on gRPC server.
	// This allows runn to dynamically discover the service
	reflection.Register(s)

	go func() {
		if err := s.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			panic(fmt.Sprintf("failed to serve: %v", err))
		}
	}()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	return s, lis.Addr().String(), nil
}