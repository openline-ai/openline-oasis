package service

import (
	"context"
	"fmt"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"openline-ai/oasis-api/hub"
	"openline-ai/oasis-api/proto"
	"openline-ai/oasis-api/routes"
	"openline-ai/oasis-api/test_utils"
	"openline-ai/oasis-api/util"
	"testing"
	"time"
)

var feedHub *hub.FeedHub
var messageHub *hub.MessageHub

var dft util.DialFactory
var oasisApi *OasisApiService

func init() {
	dft = test_utils.MakeDialFactoryTest()
}

func oasisApiDialer() (*grpc.ClientConn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	proto.RegisterOasisApiServiceServer(server, oasisApi)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	dialFunc := func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
	ctx := context.Background()
	return grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialFunc))
}

func setup(t *testing.T) {

	fh := hub.NewFeedHub()
	go fh.RunFeedHub(60 * time.Second)
	feedHub = fh

	mh := hub.NewMessageHub()
	go mh.RunMessageHub(60 * time.Second)
	messageHub = mh

	test_utils.SetupWebSocketServer(fh, mh, routes.AddWebSocketRoutes)
	oasisApi = NewOasisApiService(dft, fh, mh)

	t.Cleanup(func() {
		mh.MessageBroadcast <- hub.MessageItem{Id: "quit"}
		_ = <-mh.MessageBroadcast
		fh.FeedBroadcast <- hub.MessageFeed{ContactId: "quit"}
		_ = <-fh.FeedBroadcast

	})
}

func TestMessageEvent(t *testing.T) {
	id1 := int64(7)
	gabyId := int64(2)
	gabbyIdStr := fmt.Sprintf("%d", gabyId)

	test_utils.SetMessageStoreCallbacks(&test_utils.MockMessageServiceCallbacks{GetMessage: func(ctx context.Context, message *msProto.Message) (*msProto.Message, error) {
		if !assert.Equal(t, id1, message.GetId()) {
			return nil, status.Error(404, "Unexpected Message Id")
		}
		mi := &msProto.Message{
			Type:      msProto.MessageType_MESSAGE, //MESSAGE
			Message:   "Hello Gabi",
			Direction: msProto.MessageDirection_INBOUND, // INBOUND
			Channel:   msProto.MessageChannel_MAIL,      // Mail
			Username:  "gabi@example.org",
			Id:        &id1,
			Time:      timestamppb.New(time.Now()),
			Contact:   &msProto.Contact{Id: &gabyId, FirstName: "Gabriel", LastName: "Gontariu", ContactId: "77775553"},
		}
		return mi, nil
	},
		GetFeed: func(ctx context.Context, contact *msProto.Contact) (*msProto.Contact, error) {
			log.Printf("Entering GetFeed: %v", contact)
			if !assert.NotNil(t, contact.Id, "we expected the id from the path to be passed") {
				return nil, status.Error(400, "Feed ID should not be null")
			}

			if *contact.Id == gabyId {
				return &msProto.Contact{ContactId: "77775553", FirstName: "Gabriel", LastName: "Gontariu", Id: &gabyId}, nil
			}
			return nil, status.Errorf(codes.Unknown, "Error Finding Feed")
		},
	})

	setup(t)
	s := test_utils.NewWSServer(t)
	defer s.Close()

	ws1 := test_utils.MakeWSConnection(t, s, "/ws/"+gabbyIdStr)
	defer ws1.Close()

	ws2 := test_utils.MakeWSConnection(t, s, "/ws/"+gabbyIdStr)
	defer ws2.Close()

	feed_ws1 := test_utils.MakeWSConnection(t, s, "/ws")
	defer feed_ws1.Close()

	feed_ws2 := test_utils.MakeWSConnection(t, s, "/ws")
	defer feed_ws2.Close()

	conn, err := oasisApiDialer()
	if err != nil {
		t.Fatal("Unable to connect to the api!")
	}
	client := proto.NewOasisApiServiceClient(conn)
	ctx := context.Background()
	_, err = client.NewMessageEvent(ctx, &proto.OasisMessageId{
		MessageId: id1,
	})
	if err != nil {
		t.Fatal("Error sending new message event!")
	}

	var mhResponse hub.MessageItem
	err = ws1.ReadJSON(&mhResponse)
	if err != nil {
		t.Fatal("Error getting Message Event!")
	}
	assert.Equal(t, fmt.Sprintf("%d", id1), mhResponse.Id)
	assert.Equal(t, "gabi@example.org", mhResponse.Username)
	assert.Equal(t, gabbyIdStr, mhResponse.FeedId)
	assert.Equal(t, "INBOUND", mhResponse.Direction)
	assert.Equal(t, "MAIL", mhResponse.Channel)

	err = ws2.ReadJSON(&mhResponse)
	if err != nil {
		t.Fatal("Error getting Message Event!")
	}
	assert.Equal(t, fmt.Sprintf("%d", id1), mhResponse.Id)
	assert.Equal(t, "gabi@example.org", mhResponse.Username)
	assert.Equal(t, gabbyIdStr, mhResponse.FeedId)
	assert.Equal(t, "INBOUND", mhResponse.Direction)
	assert.Equal(t, "MAIL", mhResponse.Channel)

	var fhResponse hub.MessageFeed
	err = feed_ws1.ReadJSON(&fhResponse)
	if err != nil {
		t.Fatal("Error getting Feed Event!")
	}
	assert.Equal(t, "77775553", fhResponse.ContactId)
	assert.Equal(t, "Gabriel", fhResponse.FirstName)
	assert.Equal(t, "Gontariu", fhResponse.LastName)

	err = feed_ws2.ReadJSON(&fhResponse)
	if err != nil {
		t.Fatal("Error getting Feed Event!")
	}
	assert.Equal(t, "77775553", fhResponse.ContactId)
	assert.Equal(t, "Gabriel", fhResponse.FirstName)
	assert.Equal(t, "Gontariu", fhResponse.LastName)

}
