package routes

import (
	"encoding/json"
	"fmt"
	"github.com/DusanKasan/parsemail"
	"github.com/gin-gonic/gin"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	ms "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
	o "github.com/openline-ai/openline-oasis/packages/server/oasis-api/proto/generated"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
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
	MessageId string   `json:"messageId"`
	Html      string   `json:"html"`
	Subject   string   `json:"subject"`
	From      string   `json:"from"`
	To        []string `json:"to"`
	Cc        []string `json:"cc"`
	Bcc       []string `json:"bcc"`
	InReplyTo []string `json:"InReplyTo"`
	Reference []string `json:"Reference"`
}

func addMailRoutes(conf *c.Config, df util.DialFactory, rg *gin.RouterGroup) {
	mailGroup := rg.Group("/mail")
	mailGroup.POST("/fwd/", func(c *gin.Context) {
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

		mailReader := strings.NewReader(req.RawMessage)
		email, err := parsemail.Parse(mailReader) // returns Email struct and error
		if err != nil {
			log.Printf("Unable to parse Email: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("Unable to parse Email: %v", err.Error()),
			})
			return
		}

		if len(email.From) != 1 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("Email has more than one From: %v", email.From),
			})
			return
		}

		fromAddress := email.From[0].Address
		emailContent := EmailContent{
			MessageId: ensureRfcId(email.MessageID),
			Subject:   email.Subject,
			Html:      email.HTMLBody,
			From:      fromAddress,
			To:        toStringArr(email.To),
			Cc:        toStringArr(email.Cc),
			Bcc:       toStringArr(email.Bcc),
			InReplyTo: ensureRfcIds(email.InReplyTo),
			Reference: ensureRfcIds(email.References),
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
		refSize := len(email.References)
		threadId := ""
		if refSize > 0 {
			threadId = email.References[0]
		} else {
			threadId = email.MessageID
		}

		message := &ms.InputMessage{
			Type:                    ms.MessageType_EMAIL,
			Subtype:                 ms.MessageSubtype_MESSAGE,
			Content:                 &jsonContentString,
			Direction:               ms.MessageDirection_INBOUND,
			InitiatorIdentifier:     &fromAddress,
			ThreadId:                &threadId,
			ParticipantsIdentifiers: toStringArr(append(email.To, email.Cc...)),
		}
		//Store the message in message store
		msConn := util.GetMessageStoreConnection(c, df)
		defer util.CloseMessageStoreConnection(msConn)
		msClient := ms.NewMessageStoreServiceClient(msConn)

		ctx := context.Background()
		ctx = metadata.AppendToOutgoingContext(ctx, service.ApiKeyHeader, conf.Service.MessageStoreApiKey)
		ctx = metadata.AppendToOutgoingContext(ctx, service.UsernameHeader, email.To[0].Address)

		savedMessage, err := msClient.SaveMessage(ctx, message)
		if err != nil {
			se, _ := status.FromError(err)
			log.Printf("failed creating message item: status=%s message=%s", se.Code(), se.Message())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("failed creating message item: status=%s message=%s", se.Code(), se.Message()),
			})
			return
		}
		log.Printf("message item created with id: %s", savedMessage.GetConversationEventId())

		//Set up a connection to the oasis-api server.
		oasisConn := GetOasisClient(c, df)
		defer closeOasisConnection(oasisConn)
		oasisClient := o.NewOasisApiServiceClient(oasisConn)
		_, mEventErr := oasisClient.NewMessageEvent(ctx, &o.NewMessage{ConversationId: savedMessage.ConversationId, ConversationItemId: savedMessage.GetConversationEventId()})
		if mEventErr != nil {
			se, _ := status.FromError(mEventErr)
			log.Printf("failed new message event: status=%s message=%s", se.Code(), se.Message())
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprintf("message item created with id: %s", savedMessage.GetConversationEventId()),
		})
	})
}

func ensureRfcId(id string) string {
	if !strings.HasPrefix(id, "<") {
		id = fmt.Sprintf("<%s>", id)
	}
	return id
}

func ensureRfcIds(to []string) []string {
	var result []string
	for _, id := range to {
		result = append(result, ensureRfcId(id))
	}
	return result
}

func toStringArr(from []*mail.Address) []string {
	var to []string
	for _, a := range from {
		to = append(to, a.Address)
	}
	return to
}
