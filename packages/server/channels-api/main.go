package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"google.golang.org/grpc"
	"log"
	"net"
	c "openline-ai/channels-api/config"
	"openline-ai/channels-api/ent/proto"
	"openline-ai/channels-api/hub"
	"openline-ai/channels-api/routes"
	"openline-ai/channels-api/service"
)

func main() {
	conf := c.Config{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("missing required config")
		return
	}

	mh := hub.NewWebChatMessageHub()
	go mh.RunWebChatMessageHub()
	// Our server will live in the routes package
	go routes.Run(&conf, mh) // run this as a backround goroutine

	// Initialize the generated User service.
	svc := service.NewSendMessageService(&conf, mh)

	log.Printf("Attempting to start GRPC server")
	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()

	// Register the Message Item service with the server.
	proto.RegisterMessageEventServiceServer(server, svc)

	// Open port for listening to traffic.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Service.GRPCPort))
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	} else {
		log.Printf("server started on: %s", fmt.Sprintf(":%d", conf.Service.GRPCPort))
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}

}
