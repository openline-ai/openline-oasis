package routes

import (
	"fmt"
	"github.com/DusanKasan/parsemail"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	c "openline-ai/channels-api/config"
	pb "openline-ai/message-store/ent/proto"
	"strings"
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
			log.Fatalf("unable to parse json: %v", err)
			// DO SOMETHING WITH THE ERROR
		}
		c.JSON(http.StatusOK, "Mail POST endpoint. req sent: sender "+req.Sender+"; raw message: "+req.RawMessage)

		log.Printf("Got message from %s", req.Sender)
		mailReader := strings.NewReader(req.RawMessage)
		email, err := parsemail.Parse(mailReader) // returns Email struct and error
		if err != nil {
			log.Fatalf("Unable to parse Email: %v", err)
		}
		// Contact the server and print out its response.
		mi := &pb.Message{
			Type:      pb.MessageType_MESSAGE,
			Message:   email.TextBody,
			Direction: pb.MessageDirection_INBOUND,
			Channel:   pb.MessageChannel_MAIL,
			Username:  req.Sender,
		}
		//Set up a connection to the server.
		conn, err := grpc.Dial(conf.Service.MessageStore, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := pb.NewMessageStoreServiceClient(conn)

		ctx := context.Background()

		message, err := client.SaveMessage(ctx, mi)

		if err != nil {
			se, _ := status.FromError(err)
			log.Fatalf("failed creating message item: status=%s message=%s", se.Code(), se.Message())
		}
		log.Printf("message item created with id: %d", *message.Id)

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint("created"),
		})
	})
}
