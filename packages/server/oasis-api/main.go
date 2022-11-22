package main

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"google.golang.org/grpc"
	"log"
	"net"
	"openline-ai/oasis-api/hub"
	"openline-ai/oasis-api/proto"
	"openline-ai/oasis-api/util"

	c "openline-ai/oasis-api/config"
	"openline-ai/oasis-api/routes"
	"openline-ai/oasis-api/service"
)

func main() {
	flag.Parse()
	conf := c.Config{}

	if err := env.Parse(&conf); err != nil {
		fmt.Printf("missing required config")
		return
	}
	fh := hub.NewFeedHub()
	go fh.RunFeedHub()

	mh := hub.NewMessageHub()
	go mh.RunMessageHub()

	// Our server will live in the routes package
	go routes.Run(&conf, fh, mh) // run this as a background goroutine

	// Initialize the generated User service.
	df := util.MakeDialFactory(&conf)
	svc := service.NewOasisApiService(df, fh, mh)

	log.Printf("Attempting to start GRPC server")
	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()

	// Register the Message Item service with the server.
	proto.RegisterOasisApiServiceServer(server, svc)

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
