package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	proto "github.com/openline-ai/openline-oasis/packages/server/channels-api/proto/generated"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/routes"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/routes/chatHub"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/service"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

func main() {
	conf := loadConfiguration()

	mh := chatHub.NewHub()
	go mh.Run()

	oauthConfig := &oauth2.Config{
		ClientID:     conf.GMail.ClientId,
		ClientSecret: conf.GMail.ClientSecret,
		RedirectURL:  strings.Split(conf.GMail.RedirectUris, " ")[0],
		Scopes:       []string{gmail.GmailReadonlyScope, gmail.GmailComposeScope, "email", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  google.Endpoint.AuthURL,
			TokenURL: google.Endpoint.TokenURL,
		},
	}

	// Our server will live in the routes package
	go routes.Run(&conf, mh, oauthConfig) // run this as a backround goroutine

	// Initialize the generated User service.
	df := util.MakeDialFactory(&conf)
	svc := service.NewSendMessageService(&conf, df, oauthConfig, mh)

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

func loadConfiguration() c.Config {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARNING] Error loading .env file")
	}

	cfg := c.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}

	return cfg
}
