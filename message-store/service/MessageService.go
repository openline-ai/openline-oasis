package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"openline-ai/message-store/ent"
	"openline-ai/message-store/ent/messagefeed"
	"openline-ai/message-store/ent/messageitem"
	"openline-ai/message-store/ent/proto"
	pb "openline-ai/message-store/ent/proto"
	time "time"
)

type messageService struct {
	proto.UnimplementedMessageStoreServiceServer
	client *ent.Client
}

func encodeChannel(channel pb.MessageChannel) messageitem.Channel {
	switch channel {
	case pb.MessageChannel_WIDGET:
		return messageitem.ChannelCHAT
	case pb.MessageChannel_MAIL:
		return messageitem.ChannelMAIL
	case pb.MessageChannel_WHATSAPP:
		return messageitem.ChannelWHATSAPP
	case pb.MessageChannel_FACEBOOK:
		return messageitem.ChannelFACEBOOK
	case pb.MessageChannel_TWITTER:
		return messageitem.ChannelTWITTER
	case pb.MessageChannel_VOICE:
		return messageitem.ChannelVOICE
	default:
		return messageitem.ChannelCHAT
	}
}

func encodeDirection(direction pb.MessageDirection) messageitem.Direction {
	switch direction {
	case pb.MessageDirection_INBOUND:
		return messageitem.DirectionINBOUND
	case pb.MessageDirection_OUTBOUND:
		return messageitem.DirectionOUTBOUND
	default:
		return messageitem.DirectionOUTBOUND
	}
}

func encodeType(t pb.MessageType) messageitem.Type {
	switch t {
	case pb.MessageType_MESSAGE:
		return messageitem.TypeMESSAGE
	case pb.MessageType_FILE:
		return messageitem.TypeFILE
	default:
		return messageitem.TypeMESSAGE
	}
}

func decodeType(t messageitem.Type) pb.MessageType {
	switch t {
	case messageitem.TypeMESSAGE:
		return pb.MessageType_MESSAGE
	case messageitem.TypeFILE:
		return pb.MessageType_FILE
	default:
		return pb.MessageType_MESSAGE
	}
}

func decodeDirection(direction messageitem.Direction) pb.MessageDirection {
	switch direction {
	case messageitem.DirectionINBOUND:
		return pb.MessageDirection_INBOUND
	case messageitem.DirectionOUTBOUND:
		return pb.MessageDirection_OUTBOUND
	default:
		return pb.MessageDirection_OUTBOUND
	}
}

func decodeChannel(channel messageitem.Channel) pb.MessageChannel {
	switch channel {
	case messageitem.ChannelCHAT:
		return pb.MessageChannel_WIDGET
	case messageitem.ChannelMAIL:
		return pb.MessageChannel_MAIL
	case messageitem.ChannelWHATSAPP:
		return pb.MessageChannel_WHATSAPP
	case messageitem.ChannelFACEBOOK:
		return pb.MessageChannel_FACEBOOK
	case messageitem.ChannelTWITTER:
		return pb.MessageChannel_TWITTER
	case messageitem.ChannelVOICE:
		return pb.MessageChannel_VOICE
	default:
		return pb.MessageChannel_WIDGET
	}
}
func (s *messageService) SaveMessage(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	var contact string
	if message.Contact == nil {
		contact = message.Username // TODO: resolve address to contact
	} else {
		contact = message.Contact.Username
	}
	feed, err := s.client.MessageFeed.
		Create().
		SetUsername(contact).
		Save(ctx)

	if err != nil {
		se, _ := status.FromError(err)
		if se.Code() != codes.Unknown {
			return nil, status.Errorf(se.Code(), "Error upserting Feed")
		} else {
			feed, err = s.client.MessageFeed.
				Query().
				Where(messagefeed.Username(contact)).
				First(ctx)
			if err != nil {
				se, _ = status.FromError(err)
				return nil, status.Errorf(se.Code(), "Error getting existing Feed")
			}

		}
	}

	var time *time.Time = nil
	if message.GetTime() != nil {
		var timeref = message.GetTime().AsTime()
		time = &timeref
	}
	msg, err := s.client.MessageItem.
		Create().
		SetMessage(message.GetMessage()).
		SetMessageFeed(feed).
		SetChannel(encodeChannel(message.GetChannel())).
		SetNillableTime(time).
		SetUsername(message.GetUsername()).
		SetDirection(encodeDirection(message.GetDirection())).
		SetType(encodeType(message.GetType())).
		Save(ctx)

	if err != nil {
		se, _ := status.FromError(err)
		return nil, status.Errorf(se.Code(), "Error inserting message")
	}

	var id int64 = int64(msg.ID)
	mi := &pb.Message{
		Type:      decodeType(msg.Type),
		Message:   msg.Message,
		Direction: decodeDirection(msg.Direction),
		Channel:   decodeChannel(msg.Channel),
		Username:  msg.Username,
		Id:        &id,
		Contact:   &pb.Contact{Username: contact},
	}
	return mi, nil
}

func (s *messageService) GetMessages(ctx context.Context, contact *pb.Contact) (*pb.MessageList, error) {
	mf := &ent.MessageFeed{Username: contact.Username}
	ml := &pb.MessageList{
		Message: []*pb.Message,
	}
	messages, err := s.client.MessageFeed.QueryMessageItem(mf).All(ctx)
	if err != nil {
		se, _ := status.FromError(err)
		return nil, status.Errorf(se.Code(), "Error getting messages")
	}

	for _, message := range messages {
		var id int64 = int64(message.ID)
		mi := &pb.Message{
			Type:      decodeType(message.Type),
			Message:   message.Message,
			Direction: decodeDirection(message.Direction),
			Channel:   decodeChannel(message.Channel),
			Username:  message.Username,
			Id:        &id,
			Contact:   &pb.Contact{Username: contact.Username},
		}
		ml.Message = append(ml.Message, mi)
	}
	return ml, nil
}

func NewMessageService(client *ent.Client) *messageService {
	ms := new(messageService)
	ms.client = client
	return ms
}
