package service

import (
	"context"
	"fmt"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/proto"
	"log"
	op "openline-ai/oasis-api/proto"
	"openline-ai/oasis-api/routes/FeedHub"
	"openline-ai/oasis-api/routes/MessageHub"
	"openline-ai/oasis-api/util"
	"strconv"
)

type OasisApiService struct {
	op.UnimplementedOasisApiServiceServer
	df util.DialFactory
	mh *MessageHub.MessageHub
	fh *FeedHub.FeedHub
}

func (s OasisApiService) NewMessageEvent(c context.Context, oasisId *op.OasisMessageId) (*op.OasisEmpty, error) {
	conn, err := s.df.GetMessageStoreCon()
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
	feed, err := client.GetFeed(ctx, message.Contact)
	if err != nil {
		log.Printf("Unable to connect to retrieve message feed!")
		return nil, err
	}

	time := MessageHub.Time{
		Seconds: strconv.FormatInt(message.Time.Seconds, 10),
		Nanos:   fmt.Sprint(message.Time.Nanos),
	}

	log.Printf("Got a feed of %v", feed)
	// Send a feed to hub
	messageFeed := FeedHub.MessageFeed{FirstName: feed.FirstName, LastName: feed.LastName, ContactId: feed.ContactId}
	s.fh.Broadcast <- messageFeed
	log.Printf("successfully sent new feed for %v", messageFeed)

	// Send a message to hub
	messageItem := MessageHub.MessageItem{
		Username:  message.Username,
		Id:        strconv.FormatInt(*message.Id, 10),
		FeedId:    strconv.FormatInt(*feed.Id, 10),
		Direction: message.Direction.String(),
		Message:   message.Message,
		Time:      time,
		Channel:   message.Channel.String(),
	}

	s.mh.Broadcast <- messageItem
	log.Printf("successfully sent new message for %s", message.Username)
	return &op.OasisEmpty{}, nil
}

func NewOasisApiService(df util.DialFactory, fh *FeedHub.FeedHub, mh *MessageHub.MessageHub) *OasisApiService {
	ms := new(OasisApiService)
	ms.df = df
	ms.fh = fh
	ms.mh = mh
	return ms
}
