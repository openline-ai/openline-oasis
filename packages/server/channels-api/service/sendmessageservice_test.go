package service

import (
	"context"
	"fmt"
	"github.com/DusanKasan/parsemail"
	smtpmock "github.com/mocktools/go-smtp-mock/v2"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"openline-ai/channels-api/config"
	"openline-ai/channels-api/ent/proto"
	"openline-ai/channels-api/hub"
	"openline-ai/channels-api/routes"
	"openline-ai/channels-api/test_utils"
	"strings"
	"testing"
	"time"
)

var webchatMessageHub *hub.WebChatMessageHub
var channelsApi *sendMessageService
var mockMailServer *smtpmock.Server

func channelApiDialer() (*grpc.ClientConn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	proto.RegisterMessageEventServiceServer(server, channelsApi)

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

	fh := hub.NewWebChatMessageHub()
	go fh.RunWebChatMessageHub()
	webchatMessageHub = fh
	// You can pass empty smtpmock.ConfigurationAttr{}. It means that smtpmock will use default settings
	mailServer := smtpmock.New(smtpmock.ConfigurationAttr{
		LogToStdout:       true,
		LogServerActivity: true,
	})

	// To start mailServer use Start() method
	if err := mailServer.Start(); err != nil {
		fmt.Println(err)
	}

	mockMailServer = mailServer

	conf := &config.Config{}
	conf.Mail.SMTPServerPort = mailServer.PortNumber()
	conf.Mail.SMTPFromUser = "agent@agent.secretcorp.com"
	dft := test_utils.MakeDialFactoryTest(conf)
	channelsApi = NewSendMessageService(conf, dft, fh)

	test_utils.SetupWebSocketServer(fh, routes.AddWebSocketRoutes)

	t.Cleanup(func() {
		fh.MessageBroadcast <- hub.WebChatMessageItem{Kill: true}
		_ = <-fh.MessageBroadcast
		mailServer.Stop()

	})
}

func TestChatMessageEvent(t *testing.T) {
	id1 := int64(7)
	gabbyUsername := "gabi@example.org"
	gabyId := int64(2)

	test_utils.SetMessageStoreCallbacks(&test_utils.MockMessageServiceCallbacks{GetMessage: func(ctx context.Context, message *msProto.Message) (*msProto.Message, error) {
		if !assert.Equal(t, id1, message.GetId()) {
			return nil, status.Error(404, "Unexpected Message Id")
		}
		mi := &msProto.Message{
			Type:      msProto.MessageType_MESSAGE, //MESSAGE
			Message:   "Hello Gabi",
			Direction: msProto.MessageDirection_INBOUND, // INBOUND
			Channel:   msProto.MessageChannel_WIDGET,    // Chat
			Username:  "gabi@example.org",
			Id:        &id1,
			Time:      timestamppb.New(time.Now()),
			Contact:   &msProto.Contact{Id: &gabyId, FirstName: "Gabriel", LastName: "Gontariu", ContactId: "77775553"},
		}
		return mi, nil
	},
	})
	setup(t)
	s := test_utils.NewWSServer(t)
	defer s.Close()

	ws1 := test_utils.MakeWSConnection(t, s, "/ws/"+gabbyUsername)
	defer ws1.Close()

	ws2 := test_utils.MakeWSConnection(t, s, "/ws/"+gabbyUsername)
	defer ws2.Close()

	conn, err := channelApiDialer()
	if err != nil {
		t.Fatal("Unable to connect to the api!")
	}
	client := proto.NewMessageEventServiceClient(conn)
	ctx := context.Background()

	_, err = client.SendMessageEvent(ctx, &proto.MessageId{
		MessageId: id1,
	})
	if err != nil {
		t.Fatal("Error sending new message event!")
	}

	var mhResponse hub.WebChatMessageItem
	err = ws1.ReadJSON(&mhResponse)
	if err != nil {
		t.Fatal("Error getting Message Event!")
	}
	assert.Equal(t, gabbyUsername, mhResponse.Username)
	assert.Equal(t, "Hello Gabi", mhResponse.Message)

	err = ws2.ReadJSON(&mhResponse)
	if err != nil {
		t.Fatal("Error getting Message Event!")
	}
	assert.Equal(t, gabbyUsername, mhResponse.Username)
	assert.Equal(t, "Hello Gabi", mhResponse.Message)

}

func TestMailMessageEvent(t *testing.T) {
	id1 := int64(7)
	gabbyUsername := "gabi@example.org"
	gabyId := int64(2)

	test_utils.SetMessageStoreCallbacks(&test_utils.MockMessageServiceCallbacks{GetMessage: func(ctx context.Context, message *msProto.Message) (*msProto.Message, error) {
		if !assert.Equal(t, id1, message.GetId()) {
			return nil, status.Error(404, "Unexpected Message Id")
		}
		mi := &msProto.Message{
			Type:      msProto.MessageType_MESSAGE, //MESSAGE
			Message:   "Hello Gabi",
			Direction: msProto.MessageDirection_INBOUND, // INBOUND
			Channel:   msProto.MessageChannel_MAIL,      // Chat
			Username:  gabbyUsername,
			Id:        &id1,
			Time:      timestamppb.New(time.Now()),
			Contact:   &msProto.Contact{Id: &gabyId, FirstName: "Gabriel", LastName: "Gontariu", ContactId: "77775553"},
		}
		return mi, nil
	},
	})
	setup(t)
	conn, err := channelApiDialer()
	if err != nil {
		t.Fatal("Unable to connect to the api!")
	}
	client := proto.NewMessageEventServiceClient(conn)
	ctx := context.Background()

	_, err = client.SendMessageEvent(ctx, &proto.MessageId{
		MessageId: id1,
	})
	if err != nil {
		t.Fatal("Error sending new message event!")
	}

	time.Sleep(500 * time.Millisecond)
	if !assert.Equal(t, 1, len(mockMailServer.Messages()), "Email not received!") {
		return
	}
	mailReader := strings.NewReader(mockMailServer.Messages()[0].MsgRequest())
	email, err := parsemail.Parse(mailReader)
	if err != nil {
		t.Fatal("Unable to parse received ")
	}
	assert.Equal(t, email.From[0].Address, "agent@agent.secretcorp.com")
	assert.Equal(t, email.To[0].Address, gabbyUsername)
	assert.Equal(t, strings.Trim(email.TextBody, "\r\n"), "Hello Gabi")
	//mockMailServer.Messages()[0].Data()

}
