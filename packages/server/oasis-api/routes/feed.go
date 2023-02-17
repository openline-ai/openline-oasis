package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	chProto "github.com/openline-ai/openline-oasis/packages/server/channels-api/proto/generated"
	channelRoute "github.com/openline-ai/openline-oasis/packages/server/channels-api/routes"
	c "github.com/openline-ai/openline-oasis/packages/server/oasis-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
)

type FeedPostRequest struct {
	Username    string   `json:"username"`
	Message     string   `json:"message"`
	Channel     string   `json:"channel"`
	Source      string   `json:"source"`
	Direction   string   `json:"direction"`
	Destination []string `json:"destination"`
	ReplyTo     *string  `json:"replyTo,omitempty"`
}

type FeedParticipant struct {
	EMail string `json:"email"`
}
type FeedParticipantList struct {
	Participants []FeedParticipant `json:"participants"`
}

type FeedID struct {
	ID string `uri:"id"`
}

func decodeMessageType(channel string) msProto.MessageType {
	switch channel {
	case "EMAIL":
		return msProto.MessageType_EMAIL
	case "CHAT":
		return msProto.MessageType_WEB_CHAT
	case "VOICE":
		return msProto.MessageType_VOICE
	default:
		return msProto.MessageType_WEB_CHAT
	}
}

func buildEmailJson(item *msProto.FeedItem, req FeedPostRequest, msClient msProto.MessageStoreServiceClient, ctx context.Context) string {

	emailContent := channelRoute.EmailContent{
		From:    req.Username,
		To:      req.Destination,
		Subject: "Hello from Oasis",
		Html:    req.Message,
	}

	if req.ReplyTo != nil {
		lastMsg, err := msClient.GetMessage(ctx, &msProto.MessageId{
			ConversationEventId: *req.ReplyTo,
			ConversationId:      item.Id,
		})
		if err != nil {
			log.Printf("Error getting message: %v", err)
		} else {
			lastMsgJson := &channelRoute.EmailContent{}
			err = json.Unmarshal([]byte(lastMsg.Content), lastMsgJson)
			if err != nil {
				log.Printf("Error unmarshalling last message: %v", err)
			} else {
				var references []string
				copy(references, lastMsgJson.Reference)
				references = append(references, lastMsgJson.MessageId)
				emailContent.Reference = references
				emailContent.InReplyTo = []string{lastMsgJson.MessageId}
				emailContent.Subject = lastMsgJson.Subject
			}
		}
	}
	jsonContent, _ := json.Marshal(emailContent)
	return string(jsonContent)
}

func addFeedRoutes(rg *gin.RouterGroup, conf *c.Config, df util.DialFactory) {

	rg.GET("/feed", func(c *gin.Context) {
		onlyContacts := c.Query("onlyContacts")
		println(onlyContacts)
		msConn := util.GetMessageStoreConnection(c, df)
		defer util.CloseMessageStoreConnection(msConn)
		msClient := msProto.NewMessageStoreServiceClient(msConn)

		ctx := context.Background()
		ctx = metadata.AppendToOutgoingContext(ctx, service.ApiKeyHeader, conf.Service.MessageStoreApiKey)
		ctx = metadata.AppendToOutgoingContext(ctx, service.UsernameHeader, c.GetHeader(service.UsernameHeader))
		ctx = metadata.AppendToOutgoingContext(ctx, "X-Openline-IDENTITY-ID", c.GetHeader("X-Openline-IDENTITY-ID"))

		pagedRequest := &msProto.GetFeedsPagedRequest{OnlyContacts: onlyContacts == "true"}
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
		ctx = metadata.AppendToOutgoingContext(ctx, "X-Openline-IDENTITY-ID", c.GetHeader("X-Openline-IDENTITY-ID"))

		request := msProto.FeedId{Id: feedId.ID}
		feed, err := msClient.GetFeed(ctx, &request)
		log.Printf("Got the feed!")
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, feed)
	})
	rg.GET("/feed/:id/participants", func(c *gin.Context) {
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
		ctx = metadata.AppendToOutgoingContext(ctx, "X-Openline-IDENTITY-ID", c.GetHeader("X-Openline-IDENTITY-ID"))

		request := msProto.FeedId{Id: feedId.ID}

		emails, err := msClient.GetParticipants(ctx, &request)
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		response := &FeedParticipantList{}
		for _, email := range emails.GetParticipants() {
			response.Participants = append(response.Participants, FeedParticipant{EMail: email})
		}
		c.JSON(http.StatusOK, response)
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
		ctx = metadata.AppendToOutgoingContext(ctx, "X-Openline-IDENTITY-ID", c.GetHeader("X-Openline-IDENTITY-ID"))

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
		msCtx = metadata.AppendToOutgoingContext(msCtx, "X-Openline-IDENTITY-ID", c.GetHeader("X-Openline-IDENTITY-ID"))

		request := msProto.FeedId{Id: feedId.ID}
		feed, err := msClient.GetFeed(msCtx, &request)
		log.Printf("Got the feed!")
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		message := &msProto.InputMessage{
			ConversationId:      &feedId.ID,
			Type:                decodeMessageType(req.Channel),
			Subtype:             msProto.MessageSubtype_MESSAGE,
			Direction:           msProto.MessageDirection_OUTBOUND,
			InitiatorIdentifier: &req.Username,
		}
		if req.Channel == "EMAIL" {
			body := buildEmailJson(feed, req, msClient, msCtx)
			message.Content = &body
		} else {
			message.Content = &req.Message
		}
		//	message.Channel = msProto.MessageChannel_WIDGET
		//} else {
		//	message.Channel = msProto.MessageChannel_MAIL
		//}

		// inform the channel api a new message
		channelsConn := util.GetChannelsConnection(c, df)
		defer util.CloseChannelsConnection(channelsConn)
		channelsClient := chProto.NewMessageEventServiceClient(channelsConn)

		channelsCtx := context.Background()
		channelsCtx = metadata.AppendToOutgoingContext(channelsCtx, service.UsernameHeader, c.GetHeader(service.UsernameHeader))
		log.Printf("Got a header: %v", c.GetHeader("X-Openline-IDENTITY-ID"))
		channelsCtx = metadata.AppendToOutgoingContext(channelsCtx, "X-Openline-IDENTITY-ID", c.GetHeader("X-Openline-IDENTITY-ID"))

		newMsgId, err := channelsClient.SendMessageEvent(channelsCtx, message)
		if err != nil {
			c.JSON(400, gin.H{"msg": fmt.Sprintf("failed to send request to channel api: %v", err.Error())})
			return
		}

		newMsg, err := msClient.GetMessage(msCtx, newMsgId)
		if err != nil {
			c.JSON(400, gin.H{"msg": fmt.Sprintf("failed to get message from the store: %v", err.Error())})
			return
		}
		c.JSON(http.StatusOK, newMsg)
	})
}
