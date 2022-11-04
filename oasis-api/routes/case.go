package routes

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	pb "openline-ai/message-store/ent/proto"
	c "openline-ai/oasis-api/config"
)

type CasePostRequest struct {
	Username  string `json:"username"`
	Message   string `json:"message"`
	Channel   string `json:"channel"`
	Source    string `json:"source"`
	Direction string `json:"direction"`
}

type FeedID struct {
	ID int64 `uri:"id" binding:"required"`
}

func addCaseRoutes(rg *gin.RouterGroup) {
	conf := c.Config{}
	env.Parse(&conf)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{conf.Service.CorsUrl}
	corsConfig.AllowCredentials = true

	rg.Use(cors.New(corsConfig))

	rg.GET("/case", func(c *gin.Context) {
		// Contact the server and print out its response.
		empty := &pb.Empty{}
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

		contacts, err := client.GetFeeds(ctx, empty)
		log.Printf("Got the list of contacts!")
		c.JSON(http.StatusOK, contacts)
	})
	rg.GET("/case/:id/item", func(c *gin.Context) {
		var feedId FeedID
		if err := c.ShouldBindUri(&feedId); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
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

		feed := &pb.Contact{Id: &feedId.ID, Username: ""}
		ctx := context.Background()

		messages, err := client.GetMessages(ctx, feed)
		log.Printf("Got the list of messages!")
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(http.StatusOK, messages.GetMessage())
	})
	rg.GET("/case/:id", func(c *gin.Context) {
		var feedId FeedID
		if err := c.ShouldBindUri(&feedId); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
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

		feed := &pb.Contact{Id: &feedId.ID, Username: ""}
		ctx := context.Background()

		fullFeed, err := client.GetFeed(ctx, feed)
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, fullFeed)
	})
	rg.POST("/case/:id/item", func(c *gin.Context) {
		var feedId FeedID
		var req CasePostRequest

		if err := c.ShouldBindUri(&feedId); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		log.Printf("After json bind %v", req)
		message := &pb.Message{
			Username:  req.Username,
			Message:   req.Message,
			Direction: pb.MessageDirection_OUTBOUND,
			Type:      pb.MessageType_MESSAGE,
		}
		if req.Channel == "CHAT" {
			message.Channel = pb.MessageChannel_WIDGET
		} else {
			message.Channel = pb.MessageChannel_MAIL
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

		newMsg, err := client.SaveMessage(ctx, message)
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(http.StatusOK, newMsg)

	})
}
