package service

import (
	"context"
	mail "github.com/xhit/go-simple-mail/v2"
	"google.golang.org/grpc"
	"log"
	c "openline-ai/channels-api/config"
	"openline-ai/channels-api/ent/proto"
	msProto "openline-ai/message-store/ent/proto"
	"time"
)

type sendMessageService struct {
	proto.UnimplementedMessageEventServiceServer
	conf *c.Config
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

	if msg.Channel == msProto.MessageChannel_MAIL {
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
			return nil, err
		}

		// Create email
		email := mail.NewMSG()
		email.SetFrom(s.conf.Mail.SMTPFromUser)
		email.AddTo(msg.GetUsername())
		email.SetSubject("Hello")

		email.SetBody(mail.TextPlain, msg.GetMessage())

		err = email.Send(smtpClient)
		if err != nil {
			log.Printf("Unable to connect to mail server!")
			return nil, err
		}
	}
	log.Printf("Email successfully sent to %s", msg.GetUsername())
	return &proto.EventEmpty{}, nil
}

func NewSendMessageService(c *c.Config) *sendMessageService {
	ms := new(sendMessageService)
	ms.conf = c
	return ms
}
