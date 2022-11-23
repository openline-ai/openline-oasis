package test_utils

import (
	"context"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"openline-ai/channels-api/config"
	"openline-ai/channels-api/util"
	oasisProto "openline-ai/oasis-api/proto"
)

func messageStoreDialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	msProto.RegisterMessageStoreServiceServer(server, &MockMessageService{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func oasisApiDialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	oasisProto.RegisterOasisApiServiceServer(server, &MockOasisAPIService{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

type DialFactoryTestImpl struct {
	conf *config.Config
}

func (dfi DialFactoryTestImpl) GetMessageStoreCon() (*grpc.ClientConn, error) {
	ctx := context.Background()
	return grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(messageStoreDialer()))

}
func (dfi DialFactoryTestImpl) GetOasisAPICon() (*grpc.ClientConn, error) {
	ctx := context.Background()
	return grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(oasisApiDialer()))
}

func MakeDialFactoryTest() util.DialFactory {
	dfi := new(DialFactoryTestImpl)
	return *dfi
}
