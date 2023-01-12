package service

import (
	"context"
	"fmt"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/proto/generated"
	"strconv"

	op "github.com/openline-ai/openline-oasis/packages/server/oasis-api/proto/generated"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/FeedHub"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/MessageHub"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/util"
	"log"
	//"strconv"
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
	conversation, err := client.GetFeed(ctx, &msProto.FeedId{Id: newMessage.GetConversationId()})
	if err != nil {
		log.Printf("Unable to connect to retrieve the conversation!")
		return nil, err
	}

	conversationItem, err := client.GetMessage(ctx, &msProto.MessageId{Id: newMessage.GetConversationItemId()})
	if err != nil {
		log.Printf("Unable to connect to retrieve conversation item!")
		return nil, err
	}

	time := MessageHub.Time{
		Seconds: strconv.FormatInt(conversationItem.Time.Seconds, 10),
		Nanos:   fmt.Sprint(conversationItem.Time.Nanos),
	}

	//TODO
	//log.Printf("Sending a feed of %v", conversationItem)
	//messageFeed := FeedHub.MessageFeed{FirstName: conversation.ContactFirstName, LastName: conversation.ContactLastName, ContactId: conversation.ContactId}
	//s.fh.Broadcast <- messageFeed
	//log.Printf("successfully sent new feed for %v", messageFeed)

	// Send a message to hub
	messageItem := MessageHub.MessageItem{
		Username:  conversationItem.SenderUsername,
		Id:        conversationItem.Id,
		FeedId:    conversation.Id,
		Direction: conversationItem.Direction.String(),
		Message:   conversationItem.Content,
		Time:      time,
		Channel:   "1", //TODO
	}

	s.mh.Broadcast <- messageItem
	log.Printf("successfully sent new message for %s", conversationItem.SenderUsername)
	return &op.OasisEmpty{}, nil
}

func NewOasisApiService(df util.DialFactory, fh *FeedHub.FeedHub, mh *MessageHub.MessageHub) *OasisApiService {
	ms := new(OasisApiService)
	ms.df = df
	ms.fh = fh
	ms.mh = mh
	return ms
}
