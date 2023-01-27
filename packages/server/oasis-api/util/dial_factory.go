package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/config"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type DialFactory interface {
	GetMessageStoreCon() (*grpc.ClientConn, error)
	GetChannelsAPICon() (*grpc.ClientConn, error)
}

type DialFactoryImpl struct {
	conf *config.Config
}

func (dfi DialFactoryImpl) GetMessageStoreCon() (*grpc.ClientConn, error) {
	return grpc.Dial(dfi.conf.Service.MessageStoreUrl, grpc.WithInsecure())

}
func (dfi DialFactoryImpl) GetChannelsAPICon() (*grpc.ClientConn, error) {
	return grpc.Dial(dfi.conf.Service.ChannelsApi, grpc.WithInsecure())
}

func MakeDialFactory(conf *config.Config) DialFactory {
	dfi := new(DialFactoryImpl)
	dfi.conf = conf
	return *dfi
}

func GetMessageStoreConnection(c *gin.Context, df DialFactory) *grpc.ClientConn {
	conn, msErr := df.GetMessageStoreCon()
	if msErr != nil {
		log.Printf("did not connect: %v", msErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("did not connect: %v", msErr.Error()),
		})
	}
	return conn
}

func CloseMessageStoreConnection(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}

func GetChannelsConnection(c *gin.Context, df DialFactory) *grpc.ClientConn {
	conn, msErr := df.GetChannelsAPICon()
	if msErr != nil {
		log.Printf("did not connect: %v", msErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("did not connect: %v", msErr.Error()),
		})
	}
	return conn
}

func CloseChannelsConnection(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}
