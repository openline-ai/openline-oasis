package service

import (
	"context"
	"fmt"
	"log"
	msProto "openline-ai/message-store/ent/proto"
	c "openline-ai/oasis-api/config"
	"openline-ai/oasis-api/hub"
	op "openline-ai/oasis-api/proto"
	"strconv"
)

type OasisApiService struct {
	op.UnimplementedOasisApiServiceServer
	conf *c.Config
	fh   *hub.FeedHub
	mh   *hub.MessageHub
}

func (s OasisApiService) NewFeedEvent(c context.Context, contact *msProto.Contact) (*msProto.Empty, error) {
	// Send a feed to hub
	feed := hub.MessageFeed{Username: contact.Username}
	s.fh.FeedBroadcast <- feed
	log.Printf("successfully sent new feed for %s", feed.Username)
	return &msProto.Empty{}, nil
}

func (s OasisApiService) NewMessageEvent(c context.Context, message *op.OasisApiMessage) (*msProto.Empty, error) {

	time := hub.Time{
		Seconds: strconv.FormatInt(message.Message.Time.Seconds, 10),
		Nanos:   fmt.Sprint(message.Message.Time.Nanos),
	}
	// Send a feed to hub
	messageItem := hub.MessageItem{
		Username:  message.Contact.Username,
		Id:        strconv.FormatInt(*message.Message.Id, 10),
		FeedId:    strconv.FormatInt(*message.Contact.Id, 10),
		Direction: message.Message.Direction.String(),
		Message:   message.Message.Message,
		Time:      time,
		Channel:   message.Message.Channel.String(),
	}

	s.mh.MessageBroadcast <- messageItem
	log.Printf("successfully sent new message for %s", message.Message.Username)
	return &msProto.Empty{}, nil
}

func NewOasisApiService(c *c.Config, fh *hub.FeedHub, mh *hub.MessageHub) *OasisApiService {
	ms := new(OasisApiService)
	ms.conf = c
	ms.fh = fh
	ms.mh = mh
	return ms
}
