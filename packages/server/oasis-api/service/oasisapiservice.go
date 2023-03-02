package service

import (
	"context"
	"errors"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	op "github.com/openline-ai/openline-oasis/packages/server/oasis-api/proto/generated"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/ContactHub"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/FeedHub"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/MessageHub"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/util"
	"google.golang.org/grpc/metadata"
	"log"
	//"strconv"
)

type OasisApiService struct {
	op.UnimplementedOasisApiServiceServer
	df util.DialFactory
	mh *MessageHub.MessageHub
	fh *FeedHub.FeedHub
	ch *ContactHub.ContactHub
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

	reloadFeed := FeedHub.ReloadFeed{}
	s.fh.Broadcast <- reloadFeed
	log.Printf("successfully sent new feed for %v", reloadFeed)

	// Send a message to hub
	messageItem := MessageHub.MessageEvent{
		FeedId:  conversation.Id,
		Message: conversationItem,
	}

	s.mh.Broadcast <- messageItem

	participants, err := client.GetParticipantIds(ctx, &msProto.FeedId{Id: newMessage.GetConversationId()})
	if err != nil {
		log.Printf("Unable to connect to retrieve participants! feed:%s reason: %s", newMessage.GetConversationId(), err)
		return nil, err
	}

	for _, participant := range participants.Participants {
		log.Printf("Broadcasting to participant %s", participant.Id)
		contactItem := ContactHub.ContactEvent{
			ContactId: participant.Id,
			Message:   conversationItem,
		}

		s.ch.Broadcast <- contactItem
	}

	log.Printf("successfully sent new message for %s", conversationItem.SenderUsername)
	return &op.OasisEmpty{}, nil
}

func NewOasisApiService(df util.DialFactory, fh *FeedHub.FeedHub, mh *MessageHub.MessageHub, ch *ContactHub.ContactHub) *OasisApiService {
	ms := new(OasisApiService)
	ms.df = df
	ms.fh = fh
	ms.mh = mh
	ms.ch = ch
	return ms
}
