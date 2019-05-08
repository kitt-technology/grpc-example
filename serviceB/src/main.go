package main

import (
	"fmt"
	"github.com/kitt-technology/grpc-example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

var (
	pingService = PingService{}
)

func main() {
	lis, err := net.Listen("tcp", ":" + os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ping.RegisterPingServer(s, &pingService)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("Serving gRPC")
}

type PingService struct {}

func (p *PingService) Ping(ctx context.Context, request *ping.PingRequest) (*ping.PingResponse, error) {
	return &ping.PingResponse{Message: "pong"}, nil
}
