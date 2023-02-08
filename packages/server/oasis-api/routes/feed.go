package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	chProto "github.com/openline-ai/openline-oasis/packages/server/channels-api/proto/generated"
	c "github.com/openline-ai/openline-oasis/packages/server/oasis-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
)

type FeedPostRequest struct {
	Username  string `json:"username"`
	Message   string `json:"message"`
	Channel   string `json:"channel"`
	Source    string `json:"source"`
	Direction string `json:"direction"`
}

type FeedID struct {
	ID string `uri:"id"`
}

func addFeedRoutes(rg *gin.RouterGroup, conf *c.Config, df util.DialFactory) {

	rg.GET("/feed", func(c *gin.Context) {
		msConn := util.GetMessageStoreConnection(c, df)
		defer util.CloseMessageStoreConnection(msConn)
		msClient := msProto.NewMessageStoreServiceClient(msConn)

		ctx := context.Background()
		ctx = metadata.AppendToOutgoingContext(ctx, service.ApiKeyHeader, conf.Service.MessageStoreApiKey)
		ctx = metadata.AppendToOutgoingContext(ctx, service.UsernameHeader, c.GetHeader(service.UsernameHeader))

		pagedRequest := &msProto.GetFeedsPagedRequest{}
		feedList, err := msClient.GetFeeds(ctx, pagedRequest)
		if err != nil {
			log.Printf("did not get list of feeds: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("did not get list of feeds: %v", err.Error()),
			})
			return
		}

		marshal, _ := json.Marshal(feedList)
		log.Printf("Got a feed item of %s", marshal)
		c.JSON(http.StatusOK, feedList)
	})
	rg.GET("/feed/:id", func(c *gin.Context) {
		var feedId FeedID
		if err := c.ShouldBindUri(&feedId); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		msConn := util.GetMessageStoreConnection(c, df)
		defer util.CloseMessageStoreConnection(msConn)
		msClient := msProto.NewMessageStoreServiceClient(msConn)

		ctx := context.Background()
		ctx = metadata.AppendToOutgoingContext(ctx, service.ApiKeyHeader, conf.Service.MessageStoreApiKey)
		ctx = metadata.AppendToOutgoingContext(ctx, service.UsernameHeader, c.GetHeader(service.UsernameHeader))

		request := msProto.FeedId{Id: feedId.ID}
		feed, err := msClient.GetFeed(ctx, &request)
		log.Printf("Got the feed!")
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, feed)
	})
	rg.GET("/feed/:id/item", func(c *gin.Context) {
		var feedId FeedID
		if err := c.ShouldBindUri(&feedId); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		msConn := util.GetMessageStoreConnection(c, df)
		defer util.CloseMessageStoreConnection(msConn)
		msClient := msProto.NewMessageStoreServiceClient(msConn)

		ctx := context.Background()
		ctx = metadata.AppendToOutgoingContext(ctx, service.ApiKeyHeader, conf.Service.MessageStoreApiKey)
		ctx = metadata.AppendToOutgoingContext(ctx, service.UsernameHeader, c.GetHeader(service.UsernameHeader))

		request := msProto.FeedId{Id: feedId.ID}
		messages, err := msClient.GetMessagesForFeed(ctx, &request)
		log.Printf("Got the list of messages!")
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, messages.GetMessages())
	})
	rg.POST("/feed/:id/item", func(c *gin.Context) {
		var feedId FeedID
		var req FeedPostRequest

		if err := c.ShouldBindUri(&feedId); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		msConn := util.GetMessageStoreConnection(c, df)
		defer util.CloseMessageStoreConnection(msConn)
		msClient := msProto.NewMessageStoreServiceClient(msConn)

		msCtx := context.Background()
		msCtx = metadata.AppendToOutgoingContext(msCtx, service.ApiKeyHeader, conf.Service.MessageStoreApiKey)
		msCtx = metadata.AppendToOutgoingContext(msCtx, service.UsernameHeader, c.GetHeader(service.UsernameHeader))

		request := msProto.FeedId{Id: feedId.ID}
		_, err := msClient.GetFeed(msCtx, &request)
		log.Printf("Got the feed!")
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		message := &msProto.InputMessage{
			ConversationId:      &feedId.ID,
			Type:                msProto.MessageType_WEB_CHAT,
			Subtype:             msProto.MessageSubtype_MESSAGE,
			Content:             &req.Message,
			Direction:           msProto.MessageDirection_OUTBOUND,
			InitiatorIdentifier: &req.Username,
			SenderType:          msProto.SenderType_USER,
		}
		//if req.Channel == "CHAT" {
		//	message.Channel = msProto.MessageChannel_WIDGET
		//} else {
		//	message.Channel = msProto.MessageChannel_MAIL
		//}

		msStoreClient := msProto.NewMessageStoreServiceClient(msConn)
		newMsg, err := msStoreClient.SaveMessage(msCtx, message)
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		// inform the channel api a new message
		channelsConn := util.GetChannelsConnection(c, df)
		defer util.CloseChannelsConnection(channelsConn)
		channelsClient := chProto.NewMessageEventServiceClient(channelsConn)

		channelsCtx := context.Background()
		channelsCtx = metadata.AppendToOutgoingContext(channelsCtx, service.UsernameHeader, c.GetHeader(service.UsernameHeader))

		_, err = channelsClient.SendMessageEvent(channelsCtx, &chProto.MessageId{MessageId: newMsg.GetConversationEventId()})
		if err != nil {
			c.JSON(400, gin.H{"msg": fmt.Sprintf("failed to send request to channel api: %v", err.Error())})
			return
		}

		c.JSON(http.StatusOK, newMsg)
	})
}
