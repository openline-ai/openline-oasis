package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
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
		//Set up a connection to the server.
		conn, err := grpc.Dial("message-store-service.oasis-dev.svc.cluster.local:9013", grpc.WithInsecure())
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

		// On a separate RPC invocation, retrieve the user we saved previously.
		get, err := client.Get(ctx, &entpb.GetMessageItemRequest{
			Id: created.Id,
		})
		if err != nil {
			se, _ := status.FromError(err)
			log.Fatalf("failed retrieving message item: status=%s message=%s", se.Code(), se.Message())
		}
		log.Printf("retrieved message item with id=%d: %v", get.Id, get)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint("OK"),
		})
	})
}

//func createClient() entpb.MessageItemServiceClient {
//
//	return client
//}
