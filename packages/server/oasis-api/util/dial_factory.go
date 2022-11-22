package util

import (
	"google.golang.org/grpc"
	"openline-ai/oasis-api/config"
)

type DialFactory interface {
	GetMessageStoreCon() (*grpc.ClientConn, error)
	GetChannelsAPICon() (*grpc.ClientConn, error)
}

type DialFactoryImpl struct {
	conf *config.Config
}

func (dfi DialFactoryImpl) GetMessageStoreCon() (*grpc.ClientConn, error) {
	return grpc.Dial(dfi.conf.Service.MessageStore, grpc.WithInsecure())

}
func (dfi DialFactoryImpl) GetChannelsAPICon() (*grpc.ClientConn, error) {
	return grpc.Dial(dfi.conf.Service.ChannelsApi, grpc.WithInsecure())
}

func MakeDialFactory(conf *config.Config) DialFactory {
	dfi := new(DialFactoryImpl)
	dfi.conf = conf
	return *dfi
}
