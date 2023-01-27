package service

import (
	"context"
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/proto/generated"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	proto "github.com/openline-ai/openline-oasis/packages/server/channels-api/proto/generated"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/routes/chatHub"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
	"google.golang.org/grpc/metadata"
	"log"
)

type sendMessageService struct {
	proto.UnimplementedMessageEventServiceServer
	conf *c.Config
	mh   *chatHub.Hub
	df   util.DialFactory
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
		//TODO
		//mailErr := s.sendMail(msg)
		//if mailErr != nil {
		//	return nil, mailErr
		//}
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

//func (s sendMessageService) sendMail(msg *msProto.Message) error {
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

func NewSendMessageService(c *c.Config, df util.DialFactory, mh *chatHub.Hub) *sendMessageService {
	ms := new(sendMessageService)
	ms.conf = c
	ms.mh = mh
	ms.df = df
	return ms
}
