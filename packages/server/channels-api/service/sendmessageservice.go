package service

import (
	c "openline-ai/channels-api/config"
	proto "openline-ai/channels-api/proto/generated"
	"openline-ai/channels-api/routes/chatHub"
	"openline-ai/channels-api/util"
)

type sendMessageService struct {
	proto.UnimplementedMessageEventServiceServer
	conf *c.Config
	mh   *chatHub.Hub
	df   util.DialFactory
}

//func (s sendMessageService) SendMessageEvent(c context.Context, msgId *msProto.MessageId) (*msProto.EventEmpty, error) {
//conn, err := s.df.GetMessageStoreCon()
//if err != nil {
//	log.Printf("Unable to connect to message store!")
//	return nil, err
//}
//defer conn.Close()
//client := msProto.NewMessageStoreServiceClient(conn)
//
//ctx := context.Background()
//msg, err := client.GetMessage(ctx, &msProto.Id{Id: msgId.MessageId})
//if err != nil {
//	log.Printf("Unable to connect to retrieve message!")
//	return nil, err
//}
//switch msg.Channel {
//case msProto.MessageChannel_MAIL:
//	mailErr := s.sendMail(msg)
//	if mailErr != nil {
//		return nil, mailErr
//	}
//	return &proto.EventEmpty{}, nil
//case msProto.MessageChannel_WIDGET:
//	webChatErr := s.sendWebChat(msg)
//	if webChatErr != nil {
//		return nil, webChatErr
//	}
//	return &proto.EventEmpty{}, nil
//default:
//	err := fmt.Errorf("unknown channel: %s", msg.Channel)
//	return nil, err
//}
//}

//func (s sendMessageService) sendWebChat(msg *msProto.Message) error {
//	// Send a message to the hub
//	messageItem := chatHub.MessageItem{
//		Username: *msg.Username,
//		Message:  msg.Message,
//	}
//
//	s.mh.Broadcast <- messageItem
//	log.Printf("successfully sent new message for %s", *msg.Username)
//	return nil
//}

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
