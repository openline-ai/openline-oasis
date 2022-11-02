package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"openline-ai/message-store/ent"
	"openline-ai/message-store/ent/proto/entpb"

	c "openline-ai/message-store/config"

	"github.com/caarlos0/env/v6"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9009, "The grpc server port")
)

func main() {
	conf := c.Config{}
	env.Parse(&conf)
	var connUrl = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Name, conf.DB.Password)
	log.Printf("Connecting to database %s", connUrl)
	client, err := ent.Open("postgres", connUrl)
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

	// Register the Message Item service with the server.
	entpb.RegisterMessageItemServiceServer(server, svc)

	// Open port for listening to traffic.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	} else {
		log.Printf("server started on: %s", fmt.Sprintf(":%d", *port))
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
