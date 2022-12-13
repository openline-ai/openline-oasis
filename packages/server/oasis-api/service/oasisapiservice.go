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

func (s OasisApiService) NewMessageEvent(c context.Context, newMessage *op.NewMessage) (*op.OasisEmpty, error) {
	conn, err := s.df.GetMessageStoreCon()
	if err != nil {
		log.Printf("Unable to connect to message store!")
		return nil, err
	}
	defer conn.Close()
	client := msProto.NewMessageStoreServiceClient(conn)

	ctx := context.Background()
	conversation, err := client.GetFeed(ctx, &msProto.Id{Id: newMessage.GetConversationId()})
	if err != nil {
		log.Printf("Unable to connect to retrieve the conversation!")
		return nil, err
	}

	conversationItem, err := client.GetMessage(ctx, &msProto.Id{Id: newMessage.GetConversationItemId()})
	if err != nil {
		log.Printf("Unable to connect to retrieve conversation item!")
		return nil, err
	}

	time := MessageHub.Time{
		Seconds: strconv.FormatInt(conversationItem.Time.Seconds, 10),
		Nanos:   fmt.Sprint(conversationItem.Time.Nanos),
	}

	log.Printf("Sending a feed of %v", conversationItem)
	// Send a feed to hub
	messageFeed := FeedHub.MessageFeed{FirstName: conversation.ContactFirstName, LastName: conversation.ContactLastName, ContactId: conversation.ContactId}
	s.fh.Broadcast <- messageFeed
	log.Printf("successfully sent new feed for %v", messageFeed)

	// Send a message to hub
	messageItem := MessageHub.MessageItem{
		Username:  *conversationItem.Username,
		Id:        strconv.FormatInt(*conversationItem.Id, 10),
		FeedId:    strconv.FormatInt(conversation.Id, 10),
		Direction: conversationItem.Direction.String(),
		Message:   conversationItem.Message,
		Time:      time,
		Channel:   conversationItem.Channel.String(),
	}

	s.mh.Broadcast <- messageItem
	log.Printf("successfully sent new message for %s", *conversationItem.Username)
	return &op.OasisEmpty{}, nil
}

func NewOasisApiService(df util.DialFactory, fh *FeedHub.FeedHub, mh *MessageHub.MessageHub) *OasisApiService {
	ms := new(OasisApiService)
	ms.df = df
	ms.fh = fh
	ms.mh = mh
	return ms
}
