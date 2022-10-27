package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"openline-ai/message-store/ent/proto/entpb"
)

type MailPostRequest struct {
	Sender  string
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
		c.JSON(http.StatusOK, "Mail POST endpoint. req sent: sender "+req.Sender+"; raw message: "+req.RawMail)

		// Contact the server and print out its response.
		mi := &entpb.MessageItem{
			Type:      entpb.MessageItem_MESSAGE,
			Username:  req.Sender,
			Message:   req.RawMail,
			Direction: entpb.MessageItem_INBOUND,
			Channel:   entpb.MessageItem_MAIL,
		}
		res, err := client.Create(c, &entpb.CreateMessageItemRequest{MessageItem: mi})

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

func createClient() entpb.MessageItemServiceClient {
	//Set up a connection to the server.
	conn, err := grpc.Dial("message-store-service.openline-development.svc.cluster.local:9009", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := entpb.NewMessageItemServiceClient(conn)
	return client
}
