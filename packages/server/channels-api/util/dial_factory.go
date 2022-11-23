package util

import (
	"google.golang.org/grpc"
	"openline-ai/channels-api/config"
)

type DialFactory interface {
	GetMessageStoreCon() (*grpc.ClientConn, error)
	GetOasisAPICon() (*grpc.ClientConn, error)
}

type DialFactoryImpl struct {
	conf *config.Config
}

func (dfi DialFactoryImpl) GetMessageStoreCon() (*grpc.ClientConn, error) {
	return grpc.Dial(dfi.conf.Service.MessageStore, grpc.WithInsecure())

}
func (dfi DialFactoryImpl) GetOasisAPICon() (*grpc.ClientConn, error) {
	return grpc.Dial(dfi.conf.Service.OasisApiUrl, grpc.WithInsecure())
}

func MakeDialFactory(conf *config.Config) DialFactory {
	dfi := new(DialFactoryImpl)
	dfi.conf = conf
	return *dfi
}
