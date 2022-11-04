package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	pb "openline-ai/message-store/ent/proto"
	c "openline-ai/oasis-api/config"
)

type FeedPostRequest struct {
	Username  string `json:"username"`
	Message   string `json:"message"`
	Channel   string `json:"channel"`
	Source    string `json:"source"`
	Direction string `json:"direction"`
}

type FeedID struct {
	ID int64 `uri:"id" binding:"required"`
}

func addFeedRoutes(rg *gin.RouterGroup, conf c.Config) {

	rg.GET("/feed", func(c *gin.Context) {
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
	rg.GET("/feed/:id/item", func(c *gin.Context) {
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

		contact := &pb.Contact{Id: &feedId.ID, Username: ""}
		pageInfo := &pb.PageInfo{PageSize: 100}
		pageContact := &pb.PagedContact{Page: pageInfo, Contact: contact}
		ctx := context.Background()

		messages, err := client.GetMessages(ctx, pageContact)
		log.Printf("Got the list of messages!")
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(http.StatusOK, messages.GetMessage())
	})
	rg.GET("/feed/:id", func(c *gin.Context) {
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
	rg.POST("/feed/:id/item", func(c *gin.Context) {
		var feedId FeedID
		var req FeedPostRequest

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
