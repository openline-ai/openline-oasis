package test_utils

import (
	"context"
	chanProto "github.com/openline-ai/openline-oasis/packages/server/channels-api/proto/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MockSendMessageService struct {
	chanProto.UnimplementedMessageEventServiceServer
}

type MockChannelApi struct {
	SendMessageEvent func(context.Context, *chanProto.MessageId) (*chanProto.EventEmpty, error)
}

var channelCallbacks *MockChannelApi

func SetChannelApiCallbacks(c *MockChannelApi) {
	channelCallbacks = c
}

func (s MockSendMessageService) SendMessageEvent(c context.Context, msgId *chanProto.MessageId) (*chanProto.EventEmpty, error) {
	if channelCallbacks != nil && channelCallbacks.SendMessageEvent != nil {
		return channelCallbacks.SendMessageEvent(c, msgId)
	}
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageEvent not implemented")
}
