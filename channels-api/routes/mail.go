package routes

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	c "openline-ai/channels-api/config"
	"openline-ai/message-store/ent/proto/entpb"
)

type MailPostRequest struct {
	Sender     string
	RawMessage string
	Subject    string
	ApiKey     string
}

func addMailRoutes(rg *gin.RouterGroup) {
	conf := c.Config{}
	env.Parse(&conf)
	mail := rg.Group("/mail")
	mail.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "mail get")
	})
	mail.POST("/", func(c *gin.Context) {
		var req MailPostRequest
		if err := c.BindJSON(&req); err != nil {
			// DO SOMETHING WITH THE ERROR
		}
		c.JSON(http.StatusOK, "Mail POST endpoint. req sent: sender "+req.Sender+"; raw message: "+req.RawMessage)

		// Contact the server and print out its response.
		mi := &entpb.MessageItem{
			Type:      entpb.MessageItem_MESSAGE,
			Username:  req.Sender,
			Message:   req.RawMessage,
			Direction: entpb.MessageItem_INBOUND,
			Channel:   entpb.MessageItem_MAIL,
		}
		//Set up a connection to the server.
		conn, err := grpc.Dial(conf.Service.MessageStore, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := entpb.NewMessageItemServiceClient(conn)

		ctx := context.Background()

		created, err := client.Create(ctx, &entpb.CreateMessageItemRequest{MessageItem: mi})

		if err != nil {
			se, _ := status.FromError(err)
			log.Fatalf("failed creating message item: status=%s message=%s", se.Code(), se.Message())
		}
		log.Printf("message item created with id: %d", created.Id)

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(created),
		})
	})
}
