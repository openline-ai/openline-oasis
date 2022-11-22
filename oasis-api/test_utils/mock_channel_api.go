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

func (s MockSendMessageService) SendMessageEvent(c context.Context, msgId *chanProto.MessageId) (*chanProto.EventEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageEvent not implemented")
}
