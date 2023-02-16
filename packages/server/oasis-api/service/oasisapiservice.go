package service

import (
	"context"
	"errors"
	"fmt"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	"google.golang.org/grpc/metadata"
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
	md, ok := metadata.FromIncomingContext(c)
	if !ok {
		return nil, errors.New("unable to parse metadata from context")
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	conn, err := s.df.GetMessageStoreCon()
	if err != nil {
		log.Printf("Unable to connect to message store!")
		return nil, err
	}
	defer conn.Close()
	client := msProto.NewMessageStoreServiceClient(conn)
	conversation, err := client.GetFeed(ctx, &msProto.FeedId{Id: newMessage.GetConversationId()})
	if err != nil {
		log.Printf("Unable to connect to retrieve the conversation!")
		return nil, err
	}

	conversationItem, err := client.GetMessage(ctx, &msProto.MessageId{ConversationEventId: newMessage.GetConversationItemId()})
	if err != nil {
		log.Printf("Unable to connect to retrieve conversation item!")
		return nil, err
	}

	time := MessageHub.Time{
		Seconds: strconv.FormatInt(conversationItem.Time.Seconds, 10),
		Nanos:   fmt.Sprint(conversationItem.Time.Nanos),
	}

	reloadFeed := FeedHub.ReloadFeed{}
	s.fh.Broadcast <- reloadFeed
	log.Printf("successfully sent new feed for %v", reloadFeed)

	// Send a message to hub
	messageItem := MessageHub.MessageItem{
		Username:  conversationItem.SenderUsername,
		Id:        conversationItem.MessageId.ConversationEventId,
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
