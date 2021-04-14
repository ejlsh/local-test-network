package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ejlsh/local-test-network/cmd/ltn/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) StartNetwork(ctx context.Context, in *pb.CreateNetworkRequest) (*pb.CreateNetworkResponse, error) {
	return &pb.CreateNetworkResponse{Message: "Hello"}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNetworkServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
