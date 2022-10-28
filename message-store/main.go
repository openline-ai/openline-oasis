package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"openline-ai/message-store/ent"
	"openline-ai/message-store/ent/proto/entpb"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9013, "The grpc server port")
)

func main() {
	client, err := ent.Open("postgres", "host=oasis-postgres-service.oasis-dev.svc.cluster.local port=5432 user=openline-oasis dbname=openline-oasis password=my-secret-password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize the generated User service.
	svc := entpb.NewMessageItemService(client)

	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()

	// Register the User service with the server.
	entpb.RegisterMessageItemServiceServer(server, svc)

	// Open port for listening to traffic.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
