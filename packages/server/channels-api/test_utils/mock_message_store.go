package test_utils

import (
	"context"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MockMessageServiceCallbacks struct {
	SaveMessage func(context.Context, *msProto.Message) (*msProto.Message, error)
	GetMessage  func(context.Context, *msProto.Message) (*msProto.Message, error)
	GetMessages func(context.Context, *msProto.PagedContact) (*msProto.MessageList, error)
	GetFeeds    func(context.Context, *msProto.Empty) (*msProto.FeedList, error)
	GetFeed     func(context.Context, *msProto.Contact) (*msProto.Contact, error)
}

var messageCallbacks = &MockMessageServiceCallbacks{}

func SetMessageStoreCallbacks(c *MockMessageServiceCallbacks) {
	messageCallbacks = c
}

type MockMessageService struct {
	msProto.UnimplementedMessageStoreServiceServer
}

func (s *MockMessageService) SaveMessage(ctx context.Context, message *msProto.Message) (*msProto.Message, error) {
	if messageCallbacks.SaveMessage != nil {
		return messageCallbacks.SaveMessage(ctx, message)
	}
	return nil, status.Errorf(codes.Unimplemented, "method SaveMessage not implemented")
}

func (s *MockMessageService) GetMessage(ctx context.Context, message *msProto.Message) (*msProto.Message, error) {
	if messageCallbacks.GetMessage != nil {
		return messageCallbacks.GetMessage(ctx, message)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}

func (s *MockMessageService) GetMessages(ctx context.Context, pc *msProto.PagedContact) (*msProto.MessageList, error) {
	if messageCallbacks.GetMessages != nil {
		return messageCallbacks.GetMessages(ctx, pc)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}

func (s *MockMessageService) GetFeeds(ctx context.Context, empty *msProto.Empty) (*msProto.FeedList, error) {
	if messageCallbacks.GetFeeds != nil {
		return messageCallbacks.GetFeeds(ctx, empty)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetFeeds not implemented")
}

func (s *MockMessageService) GetFeed(ctx context.Context, contact *msProto.Contact) (*msProto.Contact, error) {
	if messageCallbacks.GetFeed != nil {
		return messageCallbacks.GetFeed(ctx, contact)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
