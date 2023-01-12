package routes

import (
	//"bytes"
	//"encoding/json"
	//"fmt"
	//"github.com/DusanKasan/parsemail"
	"github.com/gin-gonic/gin"
	//pb "github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/proto"
	//"golang.org/x/net/context"
	//"google.golang.org/grpc/status"
	//"log"
	"net/http"
	c "openline-ai/channels-api/config"
	"openline-ai/channels-api/util"
	//pbOasis "openline-ai/oasis-api/proto"
	//"strings"
)

type MailPostRequest struct {
	Sender     string `json:"sender"`
	RawMessage string `json:"rawMessage"`
	Subject    string `json:"subject"`
	ApiKey     string `json:"api-key"`
}

func addMailRoutes(conf *c.Config, df util.DialFactory, rg *gin.RouterGroup) {
	mail := rg.Group("/mail")
	mail.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "mail get")
	})
	//mail.POST("/", func(c *gin.Context) {
	//	var req MailPostRequest
	//	if err := c.BindJSON(&req); err != nil {
	//		log.Printf("unable to parse json: %v", err.Error())
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"result": fmt.Sprintf("unable to parse json: %v", err.Error()),
	//		})
	//		return
	//	}
	//
	//	if conf.Mail.ApiKey != req.ApiKey {
	//		c.JSON(http.StatusForbidden, gin.H{"result": "Invalid API Key"})
	//		return
	//	}
	//
	//	log.Printf("Got message from %s", req.Sender)
	//	mailReader := strings.NewReader(req.RawMessage)
	//	email, err := parsemail.Parse(mailReader) // returns Email struct and error
	//	if err != nil {
	//		log.Printf("Unable to parse Email: %v", err.Error())
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"result": fmt.Sprintf("Unable to parse Email: %v", err.Error()),
	//		})
	//		return
	//	}
	//	//Contact the server and print out its response.
	//	mi := &pb.Message{
	//		Type:      pb.MessageType_MESSAGE,
	//		Message:   strings.Trim(email.TextBody, "\r\n"),
	//		Direction: pb.MessageDirection_INBOUND,
	//		Channel:   pb.MessageChannel_MAIL,
	//		Username:  &req.Sender,
	//	}
	//
	//	//Set up a connection to the oasis-api server.
	//	oasisConn, oasisErr := df.GetOasisAPICon()
	//	if oasisErr != nil {
	//		log.Printf("did not connect: %v", oasisErr.Error())
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"result": fmt.Sprintf("did not connect: %v", oasisErr.Error()),
	//		})
	//		return
	//	}
	//	defer oasisConn.Close()
	//	oasisClient := pbOasis.NewOasisApiServiceClient(oasisConn)
	//
	//	//Set up a connection to the message store server.
	//	msConn, msErr := df.GetMessageStoreCon()
	//	if msErr != nil {
	//		log.Printf("did not connect: %v", msErr)
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"result": fmt.Sprintf("did not connect: %v", msErr),
	//		})
	//		return
	//	}
	//	defer msConn.Close()
	//	msClient := pb.NewMessageStoreServiceClient(msConn)
	//
	//	ctx := context.Background()
	//
	//	message, saveErr := msClient.SaveMessage(ctx, mi)
	//	if saveErr != nil {
	//		se, _ := status.FromError(saveErr)
	//		log.Printf("failed creating message item: status=%s message=%s", se.Code(), se.Message())
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"result": fmt.Sprintf("failed creating message item: status=%s message=%s", se.Code(), se.Message()),
	//		})
	//		return
	//	} else {
	//		_, mEventErr := oasisClient.NewMessageEvent(ctx, &pbOasis.NewMessage{ConversationId: *message.FeedId, ConversationItemId: *message.Id})
	//		if mEventErr != nil {
	//			se, _ := status.FromError(mEventErr)
	//			log.Printf("failed new message event: status=%s message=%s", se.Code(), se.Message())
	//		}
	//	}
	//
	//	log.Printf("message item created with id: %d", *message.Id)
	//
	//	if conf.WebChat.SlackWebhookUrl != "" {
	//		values := map[string]string{"text": fmt.Sprintf("New email arrived from: %s\n%s", *message.Username, message.Message)}
	//		json_data, _ := json.Marshal(values)
	//
	//		http.Post(conf.WebChat.SlackWebhookUrl, "application/json", bytes.NewBuffer(json_data))
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"result": fmt.Sprintf("message item created with id: %d", *message.Id),
	//	})
	//})
}
