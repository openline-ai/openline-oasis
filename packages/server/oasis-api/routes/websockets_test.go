package routes

import (
	"github.com/stretchr/testify/assert"
	"openline-ai/oasis-api/hub"
	"openline-ai/oasis-api/test_utils"
	"testing"
	"time"
)

var feedHub *hub.FeedHub
var messageHub *hub.MessageHub

func setup(t *testing.T) {

	fh := hub.NewFeedHub()
	go fh.RunFeedHub()
	feedHub = fh

	mh := hub.NewMessageHub()
	go mh.RunMessageHub()
	messageHub = mh

	test_utils.SetupWebSocketServer(fh, mh, AddWebSocketRoutes)

	t.Cleanup(func() {
		mh.MessageBroadcast <- hub.MessageItem{Id: "quit"}
		_ = <-mh.MessageBroadcast
		fh.FeedBroadcast <- hub.MessageFeed{ContactId: "quit"}
		_ = <-fh.FeedBroadcast

	})
}

func TestWebsocketCleanup(t *testing.T) {
	setup(t)
	s := test_utils.NewWSServer(t)
	defer s.Close()
	ws := test_utils.MakeWSConnection(t, s, "/ws")
	assert.Equal(t, 1, len(feedHub.Clients))
	ws.Close()
	time.Sleep(2 * time.Second)
	assert.Equal(t, 0, len(feedHub.Clients))

	ws1 := test_utils.MakeWSConnection(t, s, "/ws/1")
	assert.Equal(t, 1, len(messageHub.Clients), "incorrecct number of feeds")
	assert.Equal(t, 1, len(messageHub.Clients["1"]), "incorrecct number of messages")
	ws2 := test_utils.MakeWSConnection(t, s, "/ws/1")
	assert.Equal(t, 1, len(messageHub.Clients), "incorrecct number of feeds")
	assert.Equal(t, 2, len(messageHub.Clients["1"]), "incorrecct number of messages")
	ws3 := test_utils.MakeWSConnection(t, s, "/ws/2")
	assert.Equal(t, 2, len(messageHub.Clients), "incorrecct number of feeds")
	assert.Equal(t, 1, len(messageHub.Clients["2"]), "incorrecct number of messages")

	ws1.Close()
	time.Sleep(2 * time.Second)
	assert.Equal(t, 2, len(messageHub.Clients), "incorrecct number of feeds")
	assert.Equal(t, 1, len(messageHub.Clients["1"]), "incorrecct number of messages")

	ws2.Close()
	time.Sleep(2 * time.Second)
	assert.Equal(t, 1, len(messageHub.Clients), "incorrecct number of feeds")

	ws3.Close()
	time.Sleep(2 * time.Second)
	assert.Equal(t, 0, len(messageHub.Clients), "incorrecct number of feeds")
}
