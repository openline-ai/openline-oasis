package service

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"openline-ai/message-store/ent"
	"openline-ai/message-store/ent/proto"
	pb "openline-ai/message-store/ent/proto"
)

type messageService struct {
	proto.UnimplementedMessageStoreServiceServer
	client *ent.Client
}

func (s *messageService) SaveMessage(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	var contact;
	if message.Contact == nil {
		contact = message.Username // TODO: resolve address to contact
	} else {
		contact = message.Contact.Username;
	}
	feed, err := s.client.MessageFeed.
		Create().
		SetUsername(contact).
		Save(ctx)

	if err != nil {
		se, _ := status.FromError(err)
		if se.Code() != codes.AlreadyExists {
			return nil, status.Errorf(se.Code(), "Error upserting Feed")
		}
	}

	s.client.MessageItem

	return nil, status.Errorf(codes.Unimplemented, "method SaveMessage not implemented")
}
func (s *messageService) GetMessages(context.Context, *pb.Contact) (*pb.MessageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}

func NewMessageService(client *ent.Client) *messageService {
	ms := new(messageService)
	ms.client = client
	return ms
}
