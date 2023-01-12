package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	ms "github.com/openline-ai/openline-customer-os/packages/server/message-store/generated/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	c "openline-ai/channels-api/config"
	"openline-ai/channels-api/util"
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

		//Contact the server and print out its response.
		message := &ms.WebChatInputMessage{
			Type:       ms.MessageType_MESSAGE,
			Message:    &req.Message,
			Direction:  ms.MessageDirection_INBOUND,
			Email:      &req.Username,
			SenderType: ms.SenderType_CONTACT,
		}

		//Set up a connection to the oasis-api server.
		//oasisConn, oasisErr := df.GetOasisAPICon()
		//if oasisErr != nil {
		//	log.Printf("did not connect: %v", oasisErr)
		//	c.JSON(http.StatusInternalServerError, gin.H{
		//		"result": fmt.Sprintf("did not connect: %v", oasisErr.Error()),
		//	})
		//	return
		//}
		//defer oasisConn.Close()
		//oasisClient := pbOasis.NewOasisApiServiceClient(oasisConn)

		//Store the message in message store
		msConn := GetMessageStoreWeChatClient(c, df)
		defer closeMessageStoreConnection(msConn)
		msClient := ms.NewWebChatMessageStoreServiceClient(msConn)

		ctx := context.Background()

		savedMessage, err := msClient.SaveMessage(ctx, message)
		if err != nil {
			se, _ := status.FromError(err)
			log.Printf("failed creating message item: status=%s message=%s", se.Code(), se.Message())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("failed creating message item: status=%s message=%s", se.Code(), se.Message()),
			})
			return
		}

		//TODO send message to oasis to notify the agents
		//_, mEventErr := oasisClient.NewMessageEvent(ctx, &pbOasis.NewMessage{ConversationId: *message.FeedId, ConversationItemId: *message.Id})
		//if mEventErr != nil {
		//	se, _ := status.FromError(mEventErr)
		//	log.Printf("failed new message event: status=%s message=%s", se.Code(), se.Message())
		//}

		//log.Printf("message item created with id: %d", *message.Id)

		if conf.WebChat.SlackWebhookUrl != "" {
			values := map[string]string{"text": fmt.Sprintf("Message arrived from: %s\n%s", *message.Email, message.Message)}
			json_data, _ := json.Marshal(values)

			http.Post(conf.WebChat.SlackWebhookUrl, "application/json", bytes.NewBuffer(json_data))
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprintf("message item created with id: %s", savedMessage.Id),
		})
	})
}

func GetMessageStoreWeChatClient(c *gin.Context, df util.DialFactory) *grpc.ClientConn {
	msConn, msErr := df.GetMessageStoreCon()
	if msErr != nil {
		log.Printf("did not connect: %v", msErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("did not connect: %v", msErr.Error()),
		})
	}
	return msConn
}

func closeMessageStoreConnection(msConn *grpc.ClientConn) {
	err := msConn.Close()
	if err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}
