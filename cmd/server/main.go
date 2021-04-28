package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/ejlsh/local-test-network/cmd/proto/network"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedNetworkServiceServer
}

type NodeNotRunningError struct{}

func (m *NodeNotRunningError) Error() string {
	return "The node is not running. Slot may be 0."
}

func (s *server) RunNetwork(ctx context.Context, in *pb.RunNetworkRequest) (*pb.RunNetworkResponse, error) {
	cli := CardanoCli{}
	actions := Actions{}

	actions.CreateNetwork()

	// Make this more deterministic with tries
	time.Sleep(10 * time.Second)

	getTip, _ := cli.shelley.query.Tip()

	if getTip.Slot > 0 {
		return nil, &NodeNotRunningError{}
	}

	return &pb.RunNetworkResponse{Message: "Running Network"}, nil
}

func (s *server) RegisterStakeAddress(ctx context.Context, in *pb.RegisterStakeAddressRequest) (*pb.RegisterStakeAddressResponse, error) {
	cli := CardanoCli{}
	cli.shelley.address.keyGen("verificationKeyFileOutputPath", "signingKeyFileOutputPath")
	cli.shelley.stakeAddress.keyGen("verificationKeyFilePath", "signingKeyFilePath")
	cli.shelley.address.build("paymentVerificationKeyFilePath", "stakeVerificationKeyFilePath", "outFileOutputPath")
	cli.shelley.stakeAddress.build("stakeVerificationKeyFilePath", "outFileOutputPath")
	return &pb.RegisterStakeAddressResponse{Message: "Register Stake Address"}, nil
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
