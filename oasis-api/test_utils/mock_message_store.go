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

var callbacks = &MockMessageServiceCallbacks{}

func SetMessageStoreCallbacks(c *MockMessageServiceCallbacks) {
	callbacks = c
}

type MockMessageService struct {
	msProto.UnimplementedMessageStoreServiceServer
}

func (s *MockMessageService) SaveMessage(ctx context.Context, message *msProto.Message) (*msProto.Message, error) {
	if callbacks.SaveMessage != nil {
		return callbacks.SaveMessage(ctx, message)
	}
	return nil, status.Errorf(codes.Unimplemented, "method SaveMessage not implemented")
}

func (s *MockMessageService) GetMessage(ctx context.Context, message *msProto.Message) (*msProto.Message, error) {
	if callbacks.GetMessage != nil {
		return callbacks.GetMessage(ctx, message)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}

func (s *MockMessageService) GetMessages(ctx context.Context, pc *msProto.PagedContact) (*msProto.MessageList, error) {
	if callbacks.GetMessages != nil {
		return callbacks.GetMessages(ctx, pc)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}

func (s *MockMessageService) GetFeeds(ctx context.Context, empty *msProto.Empty) (*msProto.FeedList, error) {
	if callbacks.GetFeeds != nil {
		return callbacks.GetFeeds(ctx, empty)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetFeeds not implemented")
}

func (s *MockMessageService) GetFeed(ctx context.Context, contact *msProto.Contact) (*msProto.Contact, error) {
	if callbacks.GetFeed != nil {
		return callbacks.GetFeed(ctx, contact)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
