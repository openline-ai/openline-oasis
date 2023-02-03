package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	mimemail "github.com/emersion/go-message/mail"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	proto "github.com/openline-ai/openline-oasis/packages/server/channels-api/proto/generated"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/repository"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/repository/entity"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/routes"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/routes/chatHub"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"time"
)

type sendMessageService struct {
	proto.UnimplementedMessageEventServiceServer
	conf        *c.Config
	mh          *chatHub.Hub
	df          util.DialFactory
	repos       *repository.PostgresRepositories
	oauthConfig *oauth2.Config
}

func (s sendMessageService) SendMessageEvent(c context.Context, msgId *proto.MessageId) (*proto.EventEmpty, error) {
	username, err := service.GetUsernameMetadataForGRPC(c)
	if err != nil {
		log.Printf("Missing username header")
		return nil, err
	}

	conn, err := s.df.GetMessageStoreCon()
	if err != nil {
		log.Printf("Unable to connect to message store!")
		return nil, err
	}
	defer conn.Close()
	client := msProto.NewMessageStoreServiceClient(conn)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, service.ApiKeyHeader, s.conf.Service.MessageStoreApiKey)
	ctx = metadata.AppendToOutgoingContext(ctx, service.UsernameHeader, *username)

	msg, err := client.GetMessage(ctx, &msProto.MessageId{Id: msgId.MessageId})
	if err != nil {
		log.Printf("Unable to connect to retrieve message!")
		return nil, err
	}
	switch msg.Type {
	case msProto.MessageType_EMAIL:
		mailErr := s.sendMail(msg)
		if mailErr != nil {
			return nil, mailErr
		}
		return &proto.EventEmpty{}, nil
	case msProto.MessageType_WEB_CHAT:
		webChatErr := s.sendWebChat(msg)
		if webChatErr != nil {
			return nil, webChatErr
		}
		return &proto.EventEmpty{}, nil
	default:
		err := fmt.Errorf("unknown channel: %s", msg.Type)
		return nil, err
	}
}

func (s sendMessageService) sendWebChat(msg *msProto.Message) error {
	// Send a message to the hub
	messageItem := chatHub.MessageItem{
		Username: msg.ConversationInitiatorUsername,
		Message:  msg.Content,
	}

	s.mh.Broadcast <- messageItem
	log.Printf("successfully sent new message for %s", msg.ConversationInitiatorUsername)
	return nil
}

func (s sendMessageService) sendMail(msg *msProto.Message) error {
	tok, err := s.repos.GmailAuthTokensRepository.Get(msg.SenderUsername)
	if err != nil {
		log.Printf("Unable to get gmail auth token for %s", msg.SenderUsername)
		return err
	}

	expired := !tok.Valid()
	client := s.oauthConfig.Client(context.Background(), tok)
	if expired {
		bytes, err := json.Marshal(tok)
		if err == nil {
			s.repos.GmailAuthTokensRepository.Save(&entity.GmailAuthToken{Email: msg.SenderUsername, Token: string(bytes)})
		} else {
			log.Printf("Unable to save new token for %s", msg.SenderUsername)
		}
	}

	srv, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))

	jsonMail := &routes.EmailContent{}
	err = json.Unmarshal([]byte(msg.Content), jsonMail)
	if err != nil {
		log.Printf("Unable to parse email content for %s", msg.SenderUsername)
		return err
	}
	fromAddress := []*mimemail.Address{{"", jsonMail.From}}
	toAddress := []*mimemail.Address{{"", jsonMail.To[0]}}

	var b bytes.Buffer
	user := "me"

	// Create our mail header
	var h mimemail.Header
	h.SetDate(time.Now())
	h.SetAddressList("From", fromAddress)
	h.SetAddressList("To", toAddress)
	h.SetSubject(jsonMail.Subject)

	// Create a new mail writer
	mw, err := mimemail.CreateWriter(&b, h)
	if err != nil {
		log.Fatal(err)
	}

	// Create a text part
	tw, err := mw.CreateInline()
	if err != nil {
		log.Fatal(err)
	}
	var th mimemail.InlineHeader
	th.Set("Content-Type", "text/html")
	w, err := tw.CreatePart(th)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, jsonMail.Html)
	w.Close()
	tw.Close()

	mw.Close()

	raw := base64.StdEncoding.EncodeToString(b.Bytes())
	msgToSend := &gmail.Message{
		Raw: raw,
	}
	_, err = srv.Users.Messages.Send(user, msgToSend).Do()
	if err != nil {
		return err
	}
	return nil
}

//
//	smtpClient, err := s.df.GetSMTPClientCon()
//	if err != nil {
//		log.Printf("Unable to connect to mail server! %v", err)
//		return err
//	}
//
//	// Create email
//	email := mail.NewMSG()
//	email.SetFrom(s.conf.Mail.SMTPFromUser)
//	email.AddTo(msg.GetUsername())
//	email.SetSubject("Hello") //TODO
//
//	email.SetBody(mail.TextPlain, msg.GetMessage())
//
//	err = email.Send(smtpClient)
//	if err != nil {
//		log.Printf("Unable to send to mail server!")
//		return err
//	}
//	log.Printf("Email successfully sent to %s", msg.GetUsername())
//	return nil
//}

func NewSendMessageService(c *c.Config, df util.DialFactory, repos *repository.PostgresRepositories, oauthConfig *oauth2.Config, mh *chatHub.Hub) *sendMessageService {
	ms := new(sendMessageService)
	ms.conf = c
	ms.mh = mh
	ms.df = df
	ms.repos = repos
	ms.oauthConfig = oauthConfig
	return ms
}
