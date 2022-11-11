package service

import (
	"context"
	"fmt"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/ent/proto"
	"google.golang.org/grpc"
	"log"
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

func (s OasisApiService) NewFeedEvent(c context.Context, oasisContact *op.OasisContact) (*op.OasisEmpty, error) {
	conn, err := grpc.Dial(s.conf.Service.MessageStore, grpc.WithInsecure())
	if err != nil {
		log.Printf("Unable to connect to message store!")
		return nil, err
	}
	defer conn.Close()
	client := msProto.NewMessageStoreServiceClient(conn)

	ctx := context.Background()
	feed, err := client.GetFeed(ctx, &msProto.Contact{Username: oasisContact.Username, Id: &oasisContact.Id})
	if err != nil {
		log.Printf("Unable to connect to retrieve message!")
		return nil, err
	}
	// Send a feed to hub
	hubFeed := hub.MessageFeed{Username: feed.Username}
	s.fh.FeedBroadcast <- hubFeed
	log.Printf("successfully sent new feed for %s", feed.Username)
	return &op.OasisEmpty{}, nil
}

func (s OasisApiService) NewMessageEvent(c context.Context, oasisId *op.OasisMessageId) (*op.OasisEmpty, error) {
	conn, err := grpc.Dial(s.conf.Service.MessageStore, grpc.WithInsecure())
	if err != nil {
		log.Printf("Unable to connect to message store!")
		return nil, err
	}
	defer conn.Close()
	client := msProto.NewMessageStoreServiceClient(conn)

	ctx := context.Background()
	message, err := client.GetMessage(ctx, &msProto.Message{Id: &oasisId.MessageId})
	if err != nil {
		log.Printf("Unable to connect to retrieve message!")
		return nil, err
	}

	time := hub.Time{
		Seconds: strconv.FormatInt(message.Time.Seconds, 10),
		Nanos:   fmt.Sprint(message.Time.Nanos),
	}
	// Send a feed to hub
	messageItem := hub.MessageItem{
		Username:  message.Contact.Username,
		Id:        strconv.FormatInt(*message.Id, 10),
		FeedId:    strconv.FormatInt(*message.Contact.Id, 10),
		Direction: message.Direction.String(),
		Message:   message.Message,
		Time:      time,
		Channel:   message.Channel.String(),
	}

	s.mh.MessageBroadcast <- messageItem
	log.Printf("successfully sent new message for %s", message.Username)
	return &op.OasisEmpty{}, nil
}

func NewOasisApiService(c *c.Config, fh *hub.FeedHub, mh *hub.MessageHub) *OasisApiService {
	ms := new(OasisApiService)
	ms.conf = c
	ms.fh = fh
	ms.mh = mh
	return ms
}
