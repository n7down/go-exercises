package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/n7down/go-exersices/Grpc/internal/pb/messages"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8081, "The server port")
)

type testServer struct{}

func (t *testServer) SayHello(ctx context.Context, req *messages.HelloRequest) (*messages.HelloResponse, error) {
	return &messages.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	messages.RegisterHelloServiceServer(grpcServer, &testServer{})
	grpcServer.Serve(lis)
}
