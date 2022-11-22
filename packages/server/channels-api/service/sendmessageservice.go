package service

import (
	"context"
	"fmt"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/proto"
	mail "github.com/xhit/go-simple-mail/v2"
	"google.golang.org/grpc"
	"log"
	c "openline-ai/channels-api/config"
	"openline-ai/channels-api/ent/proto"
	"openline-ai/channels-api/hub"
	"time"
)

type sendMessageService struct {
	proto.UnimplementedMessageEventServiceServer
	conf *c.Config
	mh   *hub.WebChatMessageHub
}

func (s sendMessageService) SendMessageEvent(c context.Context, msgId *proto.MessageId) (*proto.EventEmpty, error) {
	conn, err := grpc.Dial(s.conf.Service.MessageStore, grpc.WithInsecure())
	if err != nil {
		log.Printf("Unable to connect to message store!")
		return nil, err
	}
	defer conn.Close()
	client := msProto.NewMessageStoreServiceClient(conn)

	ctx := context.Background()
	msg, err := client.GetMessage(ctx, &msProto.Message{Id: &msgId.MessageId})
	if err != nil {
		log.Printf("Unable to connect to retrieve message!")
		return nil, err
	}
	switch msg.Channel {
	case msProto.MessageChannel_MAIL:
		mailErr := sendMail(s, msg)
		if mailErr != nil {
			return nil, mailErr
		}
		return &proto.EventEmpty{}, nil
	case msProto.MessageChannel_WIDGET:
		webChatErr := sendWebChat(msg, s)
		if webChatErr != nil {
			return nil, webChatErr
		}
		return &proto.EventEmpty{}, nil
	default:
		err := fmt.Errorf("unknown channel: %s", msg.Channel)
		return nil, err
	}
}

func sendWebChat(msg *msProto.Message, s sendMessageService) error {
	// Send a message to the webchat hub
	messageItem := hub.WebChatMessageItem{
		Username: msg.Username,
		Message:  msg.Message,
	}

	s.mh.MessageBroadcast <- messageItem
	log.Printf("successfully sent new message for %s", msg.Username)
	return nil
}

func sendMail(s sendMessageService, msg *msProto.Message) error {
	server := mail.NewSMTPClient()
	server.Host = s.conf.Mail.SMTPSeverAddress
	server.Port = 465
	server.Username = s.conf.Mail.SMTPSeverUser
	server.Password = s.conf.Mail.SMTPSeverPassword
	server.Encryption = mail.EncryptionSSLTLS
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	log.Printf("Trying to connect to server %s:%d", server.Host, server.Port)
	smtpClient, err := server.Connect()
	if err != nil {
		log.Printf("Unable to connect to mail server! %v", err)
		return err
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom(s.conf.Mail.SMTPFromUser)
	email.AddTo(msg.GetUsername())
	email.SetSubject("Hello")

	email.SetBody(mail.TextPlain, msg.GetMessage())

	err = email.Send(smtpClient)
	if err != nil {
		log.Printf("Unable to send to mail server!")
		return err
	}
	log.Printf("Email successfully sent to %s", msg.GetUsername())
	return nil
}

func NewSendMessageService(c *c.Config, mh *hub.WebChatMessageHub) *sendMessageService {
	ms := new(sendMessageService)
	ms.conf = c
	ms.mh = mh
	return ms
}
