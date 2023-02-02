package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type DialFactory interface {
	GetMessageStoreCon() (*grpc.ClientConn, error)
	GetOasisAPICon() (*grpc.ClientConn, error)
}

type DialFactoryImpl struct {
	conf *config.Config
}

func (dfi DialFactoryImpl) GetMessageStoreCon() (*grpc.ClientConn, error) {
	return grpc.Dial(dfi.conf.Service.MessageStoreUrl, grpc.WithInsecure())

}
func (dfi DialFactoryImpl) GetOasisAPICon() (*grpc.ClientConn, error) {
	return grpc.Dial(dfi.conf.Service.OasisApiUrl, grpc.WithInsecure())
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
