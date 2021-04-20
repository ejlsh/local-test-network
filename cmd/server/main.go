package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ejlsh/local-test-network/cmd/proto/network"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedNetworkServiceServer
}

func (s *server) RunNetwork(ctx context.Context, in *pb.RunNetworkRequest) (*pb.RunNetworkResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.RunNetworkResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNetworkServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
