package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	ms "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
	o "github.com/openline-ai/openline-oasis/packages/server/oasis-api/proto/generated"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

type WebchatMessage struct {
	Username string `json:"username"`
	UserId   string `json:"userId"`
	Message  string `json:"message"`
}

type LoginRequest struct {
	Mail string `json:"mail"`
}

type LoginResponse struct {
	UserName  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func AddWebChatRoutes(conf *c.Config, df util.DialFactory, rg *gin.RouterGroup) {

	rg.GET("/webchat/", func(c *gin.Context) {
		email := c.Query("email")
		log.Printf(email)
		if conf.WebChat.ApiKey != c.GetHeader("WebChatApiKey") {
			c.JSON(http.StatusForbidden, gin.H{"result": "Invalid API Key"})
			return
		}
		// TODO: This needs more work. Returning email back for now
		response := LoginResponse{UserName: email, FirstName: "", LastName: ""}

		c.JSON(http.StatusOK, response)
	})

	rg.POST("/webchat/", func(c *gin.Context) {
		var req WebchatMessage

		if err := c.BindJSON(&req); err != nil {
			log.Printf("unable to parse json: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to parse json: %v", err.Error()),
			})
			return
		}

		if conf.WebChat.ApiKey != c.GetHeader("WebChatApiKey") {
			c.JSON(http.StatusForbidden, gin.H{"result": "Invalid API Key"})
			return
		}

		log.Printf("Got message from %s", req.Username)
		threadId := ""
		//Contact the server and print out its response.
		message := &ms.InputMessage{
			Type:                    ms.MessageType_WEB_CHAT,
			Subtype:                 ms.MessageSubtype_MESSAGE,
			Content:                 &req.Message,
			Direction:               ms.MessageDirection_INBOUND,
			InitiatorIdentifier:     &ms.ParticipantId{Identifier: req.Username, Type: ms.ParticipantIdType_MAILTO},
			ThreadId:                &threadId,
			ParticipantsIdentifiers: []*ms.ParticipantId{},
		}

		//Store the message in message store
		msConn := util.GetMessageStoreConnection(c, df)
		defer util.CloseMessageStoreConnection(msConn)
		msClient := ms.NewMessageStoreServiceClient(msConn)

		ctx := context.Background()
		ctx = metadata.AppendToOutgoingContext(ctx, service.ApiKeyHeader, conf.Service.MessageStoreApiKey)
		ctx = metadata.AppendToOutgoingContext(ctx, service.UsernameHeader, c.GetHeader(service.UsernameHeader))
		ctx = metadata.AppendToOutgoingContext(ctx, "X-Openline-TENANT", "openline")

		savedMessage, err := msClient.SaveMessage(ctx, message)
		if err != nil {
			se, _ := status.FromError(err)
			log.Printf("failed creating message item: status=%s message=%s", se.Code(), se.Message())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("failed creating message item: status=%s message=%s", se.Code(), se.Message()),
			})
			return
		}
		log.Printf("message item created with id: %s", savedMessage.ConversationEventId)

		//Set up a connection to the oasis-api server.
		oasisConn := GetOasisClient(c, df)
		defer closeOasisConnection(oasisConn)
		oasisClient := o.NewOasisApiServiceClient(oasisConn)
		_, mEventErr := oasisClient.NewMessageEvent(ctx, &o.NewMessage{ConversationId: savedMessage.ConversationId, ConversationItemId: savedMessage.ConversationEventId})
		if mEventErr != nil {
			se, _ := status.FromError(mEventErr)
			log.Printf("failed new message event: status=%s message=%s", se.Code(), se.Message())
		}

		if conf.WebChat.SlackWebhookUrl != "" {
			values := map[string]string{"text": fmt.Sprintf("Message arrived from: %s\n%s", message.InitiatorIdentifier.Identifier, *message.Content)}
			json_data, _ := json.Marshal(values)

			http.Post(conf.WebChat.SlackWebhookUrl, "application/json", bytes.NewBuffer(json_data))
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprintf("message item created with id: %s", savedMessage.ConversationEventId),
		})
	})
}

func GetOasisClient(c *gin.Context, df util.DialFactory) *grpc.ClientConn {
	conn, msErr := df.GetOasisAPICon()
	if msErr != nil {
		log.Printf("did not connect: %v", msErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("did not connect: %v", msErr.Error()),
		})
	}
	return conn
}

func closeOasisConnection(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}
