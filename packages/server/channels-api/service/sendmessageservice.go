package service

import (
	"context"
	"fmt"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/proto"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	c "openline-ai/channels-api/config"
	"openline-ai/channels-api/ent/proto"
	"openline-ai/channels-api/hub"
	"openline-ai/channels-api/util"
)

type sendMessageService struct {
	proto.UnimplementedMessageEventServiceServer
	conf *c.Config
	mh   *hub.WebChatMessageHub
	df   util.DialFactory
}

func (s sendMessageService) SendMessageEvent(c context.Context, msgId *proto.MessageId) (*proto.EventEmpty, error) {
	conn, err := s.df.GetMessageStoreCon()
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
		mailErr := s.sendMail(msg)
		if mailErr != nil {
			return nil, mailErr
		}
		return &proto.EventEmpty{}, nil
	case msProto.MessageChannel_WIDGET:
		webChatErr := s.sendWebChat(msg)
		if webChatErr != nil {
			return nil, webChatErr
		}
		return &proto.EventEmpty{}, nil
	default:
		err := fmt.Errorf("unknown channel: %s", msg.Channel)
		return nil, err
	}
}

func (s sendMessageService) sendWebChat(msg *msProto.Message) error {
	// Send a message to the webchat hub
	messageItem := hub.WebChatMessageItem{
		Username: msg.Username,
		Message:  msg.Message,
	}

	s.mh.MessageBroadcast <- messageItem
	log.Printf("successfully sent new message for %s", msg.Username)
	return nil
}

func (s sendMessageService) sendMail(msg *msProto.Message) error {

	smtpClient, err := s.df.GetSMTPClientCon()
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

func NewSendMessageService(c *c.Config, df util.DialFactory, mh *hub.WebChatMessageHub) *sendMessageService {
	ms := new(sendMessageService)
	ms.conf = c
	ms.mh = mh
	ms.df = df
	return ms
}
