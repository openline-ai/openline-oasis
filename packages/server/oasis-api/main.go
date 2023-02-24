package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	commonRepository "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/config"
	c "github.com/openline-ai/openline-oasis/packages/server/oasis-api/config"
	proto "github.com/openline-ai/openline-oasis/packages/server/oasis-api/proto/generated"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/FeedHub"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/MessageHub"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/service"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitDB(cfg *c.Config) (db *config.StorageDB, err error) {
	if db, err = config.NewDBConn(
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Db,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.MaxConn,
		cfg.Postgres.MaxIdleConn,
		cfg.Postgres.ConnMaxLifetime); err != nil {
		log.Fatalf("Coud not open db connection: %s", err.Error())
	}
	return
}

func main() {
	flag.Parse()
	conf := loadConfiguration()

	//GORM
	db, _ := InitDB(conf)
	defer db.SqlDB.Close()

	// NEO4J
	neo4jDriver, err := c.NewDriver(conf)
	if err != nil {
		logrus.Fatalf("Could not establish connection with neo4j at: %v, error: %v", conf.Neo4j.Target, err.Error())
	}
	ctx := context.Background()
	defer (*neo4jDriver).Close(ctx)

	commonRepositories := commonRepository.InitRepositories(db.GormDB, neo4jDriver)

	fh := FeedHub.NewFeedHub()
	go fh.Run()

	mh := MessageHub.NewMessageHub()
	go mh.Run()

	// Our server will live in the routes package
	go routes.ConfigureRoutes(conf, commonRepositories, fh, mh) // run this as a background goroutine

	// Initialize the generated User service.
	df := util.MakeDialFactory(conf)
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

func loadConfiguration() *c.Config {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARNING] Error loading .env file")
	}

	cfg := c.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}

	return &cfg
}
