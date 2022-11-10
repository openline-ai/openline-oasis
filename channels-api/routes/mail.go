package routes

import (
	"fmt"
	"github.com/DusanKasan/parsemail"
	"github.com/gin-gonic/gin"
	pb "github.com/openline-ai/openline-customer-os/packages/server/message-store/ent/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	c "openline-ai/channels-api/config"
	"strings"
)

type MailPostRequest struct {
	Sender     string `json:"sender"`
	RawMessage string `json:"rawMessage"`
	Subject    string `json:"subject"`
	ApiKey     string `json:"api-key"`
}

func addMailRoutes(conf *c.Config, rg *gin.RouterGroup) {
	mail := rg.Group("/mail")
	mail.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "mail get")
	})
	mail.POST("/", func(c *gin.Context) {
		var req MailPostRequest
		if err := c.BindJSON(&req); err != nil {
			log.Printf("unable to parse json: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to parse json: %v", err),
			})
			return
		}

		if conf.Mail.ApiKey != req.ApiKey {
			c.JSON(http.StatusForbidden, gin.H{"result": "Invalid API Key"})
			return
		}
		//c.JSON(http.StatusOK, "Mail POST endpoint. req sent: sender "+req.Sender+"; raw message: "+req.RawMessage)

		log.Printf("Got message from %s", req.Sender)
		mailReader := strings.NewReader(req.RawMessage)
		email, err := parsemail.Parse(mailReader) // returns Email struct and error
		if err != nil {
			log.Printf("Unable to parse Email: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("Unable to parse Email: %v", err),
			})
			return
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
			log.Printf("did not connect: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("did not connect: %v", err),
			})
			return
		}
		defer conn.Close()
		client := pb.NewMessageStoreServiceClient(conn)

		ctx := context.Background()

		message, err := client.SaveMessage(ctx, mi)

		if err != nil {
			se, _ := status.FromError(err)
			log.Printf("failed creating message item: status=%s message=%s", se.Code(), se.Message())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("failed creating message item: status=%s message=%s", se.Code(), se.Message()),
			})
			return
		}
		log.Printf("message item created with id: %d", *message.Id)

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprintf("message item created with id: %d", *message.Id),
		})
	})
}
