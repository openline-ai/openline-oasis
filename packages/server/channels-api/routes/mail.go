package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DusanKasan/parsemail"
	"github.com/gin-gonic/gin"
	ms "github.com/openline-ai/openline-customer-os/packages/server/message-store/proto/generated"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
	o "github.com/openline-ai/openline-oasis/packages/server/oasis-api/proto/generated"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"net/mail"
	"strings"
	//pbOasis "openline-ai/oasis-api/proto"
	//"strings"
)

type MailPostRequest struct {
	Sender     string `json:"sender"`
	RawMessage string `json:"rawMessage"`
	Subject    string `json:"subject"`
	ApiKey     string `json:"api-key"`
}

type EmailContent struct {
	Html    string   `json:"html"`
	Subject string   `json:"subject"`
	From    string   `json:"from"`
	To      []string `json:"to"`
	Cc      []string `json:"cc"`
	Bcc     []string `json:"bcc"`
}

func addMailRoutes(conf *c.Config, df util.DialFactory, rg *gin.RouterGroup) {
	mail := rg.Group("/mail")
	mail.POST("/fwd/", func(c *gin.Context) {
		var req MailPostRequest
		if err := c.BindJSON(&req); err != nil {
			log.Printf("unable to parse json: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to parse json: %v", err.Error()),
			})
			return
		}

		if conf.Mail.ApiKey != req.ApiKey {
			c.JSON(http.StatusForbidden, gin.H{"result": "Invalid API Key"})
			return
		}
		log.Printf("Mail request: %+v", req)
		log.Printf("Mail request: %+v", req.RawMessage)
		mailReader := strings.NewReader(req.RawMessage)
		email, err := parsemail.Parse(mailReader) // returns Email struct and error
		if err != nil {
			log.Printf("Unable to parse Email: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("Unable to parse Email: %v", err.Error()),
			})
			return
		}
		log.Printf("Got message from %s", email.From)
		log.Printf("Got message To %s", email.To)
		log.Printf("Text Body: %v", email.TextBody)
		log.Printf("Text Subject: %v", email.Subject)
		log.Printf("Text HTMLBody: %v", email.HTMLBody)
		log.Printf("Forwared To: %v", email.Header.Get("X-Forwarded-To"))

		if len(email.From) != 1 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("Email has more than one From: %v", email.From),
			})
			return
		}

		fromAddress := email.From[0].Address
		emailContent := EmailContent{
			Subject: email.Subject,
			Html:    email.HTMLBody,
			From:    fromAddress,
			To:      toStringArr(email.To),
			Cc:      toStringArr(email.Cc),
			Bcc:     toStringArr(email.Bcc),
		}
		jsonContent, err := json.Marshal(emailContent)
		if err != nil {
			se, _ := status.FromError(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("failed creating message content: status=%s message=%s", se.Code(), se.Message()),
			})
			return
		}
		//Contact the server and print out its response.
		jsonContentString := string(jsonContent)
		message := &ms.WebChatInputMessage{
			Type:       ms.MessageType_EMAIL,
			Subtype:    ms.MessageSubtype_MESSAGE,
			Message:    &jsonContentString,
			Direction:  ms.MessageDirection_INBOUND,
			Email:      &fromAddress,
			SenderType: ms.SenderType_CONTACT,
		}

		//Store the message in message store
		msConn := util.GetMessageStoreConnection(c, df)
		defer util.CloseMessageStoreConnection(msConn)
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
		log.Printf("message item created with id: %s", savedMessage.Id)

		//Set up a connection to the oasis-api server.
		oasisConn := GetOasisClient(c, df)
		defer closeOasisConnection(oasisConn)
		oasisClient := o.NewOasisApiServiceClient(oasisConn)
		_, mEventErr := oasisClient.NewMessageEvent(ctx, &o.NewMessage{ConversationId: savedMessage.ConversationId, ConversationItemId: savedMessage.Id})
		if mEventErr != nil {
			se, _ := status.FromError(mEventErr)
			log.Printf("failed new message event: status=%s message=%s", se.Code(), se.Message())
		}

		if conf.WebChat.SlackWebhookUrl != "" {
			values := map[string]string{"text": fmt.Sprintf("Message arrived from: %s\n%s", *message.Email, *message.Message)}
			json_data, _ := json.Marshal(values)

			http.Post(conf.WebChat.SlackWebhookUrl, "application/json", bytes.NewBuffer(json_data))
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprintf("message item created with id: %s", savedMessage.Id),
		})
	})
}

func toStringArr(from []*mail.Address) []string {
	var to []string
	for _, a := range from {
		to = append(to, a.Address)
	}
	return to
}
