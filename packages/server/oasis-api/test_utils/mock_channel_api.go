package test_utils

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	chanProto "openline-ai/channels-api/ent/proto"
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
