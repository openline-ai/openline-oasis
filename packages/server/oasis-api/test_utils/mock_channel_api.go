package test_utils

import (
	"context"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	chanProto "github.com/openline-ai/openline-oasis/packages/server/channels-api/proto/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MockSendMessageService struct {
	chanProto.UnimplementedMessageEventServiceServer
}

type MockChannelApi struct {
	SendMessageEvent func(context.Context, *msProto.InputMessage) (*msProto.Message, error)
}

var channelCallbacks *MockChannelApi

func SetChannelApiCallbacks(c *MockChannelApi) {
	channelCallbacks = c
}

func (s MockSendMessageService) SendMessageEvent(c context.Context, msg *msProto.InputMessage) (*msProto.Message, error) {
	if channelCallbacks != nil && channelCallbacks.SendMessageEvent != nil {
		return channelCallbacks.SendMessageEvent(c, msg)
	}
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageEvent not implemented")
}
