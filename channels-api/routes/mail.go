package routes

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	pb "openline-ai/oasis-common/grpc"

	"github.com/gin-gonic/gin"
)

type MailPostRequest struct {
	Senders string
	RawMail string
	Subject string
	ApiKey  string
}

func addMailRoutes(rg *gin.RouterGroup) {
	client := createClient()

	mail := rg.Group("/mail")
	mail.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "mail get")
	})
	mail.POST("/", func(c *gin.Context) {
		var req MailPostRequest
		if err := c.BindJSON(&req); err != nil {
			// DO SOMETHING WITH THE ERROR
		}
		c.JSON(http.StatusOK, "Mail POST endpoint. req sent: sender "+req.Senders+"; raw message: "+req.RawMail)

		// Contact the server and print out its response.
		omsg := &pb.OmniMessage{Type: pb.MessageType_MESSAGE,
			Username:  req.Senders,
			Direction: pb.MessageDirection_INBOUND,
			Message:   req.RawMail,
			Channel:   pb.MessageChannel_MAIL,
		}
		res, err := client.SaveMessage(c, omsg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res),
		})
	})
}

func createClient() pb.MessageStoreClient {
	// Set up a connection to the server.
	conn, err := grpc.Dial("message-store-service.openline-development.svc.cluster.local:9009", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMessageStoreClient(conn)
	return client
}
