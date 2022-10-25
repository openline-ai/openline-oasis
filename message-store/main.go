package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "openline-ai/oasis-common/grpc"
)

var (
	port = flag.Int("port", 9013, "The grpc server port")
)

type server struct {
	pb.UnimplementedMessageStoreServer
}

func (s *server) SaveMessage(ctx context.Context, in *pb.OmniMessage) (*pb.Empty, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.Empty{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageStoreServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
