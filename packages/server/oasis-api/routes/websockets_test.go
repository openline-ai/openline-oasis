package routes

import (
	"github.com/stretchr/testify/assert"
	"openline-ai/oasis-api/hub"
	"openline-ai/oasis-api/test_utils"
	"sync"
	"testing"
	"time"
)

var feedHub *hub.FeedHub
var messageHub *hub.MessageHub

func setup(t *testing.T) {

	fh := hub.NewFeedHub()
	go fh.RunFeedHub(60)
	feedHub = fh

	mh := hub.NewMessageHub()
	go mh.RunMessageHub(60)
	messageHub = mh

	test_utils.SetupWebSocketServer(fh, mh, AddWebSocketRoutes)

	t.Cleanup(func() {
		mh.MessageBroadcast <- hub.MessageItem{Id: "quit"}
		_ = <-mh.MessageBroadcast
		fh.FeedBroadcast <- hub.MessageFeed{ContactId: "quit"}
		_ = <-fh.FeedBroadcast

	})
}

func waitTimeout(t *testing.T, cond *sync.Cond, timeout time.Duration) {
	done := make(chan struct{})
	go func() {
		cond.Wait()
		close(done)
	}()
	select {
	case <-time.After(timeout):
		t.Fatal("Sync Thread Never Woke Up!")
	case <-done:
		// Wait returned
	}
}

func TestWebsocketCleanup(t *testing.T) {
	setup(t)
	s := test_utils.NewWSServer(t)
	defer s.Close()
	feedHub.Sync.L.Lock()
	ws := test_utils.MakeWSConnection(t, s, "/ws")
	waitTimeout(t, feedHub.Sync, 5*time.Second)
	feedHub.Sync.L.Unlock()

	feedHub.Sync.L.Lock()
	assert.Equal(t, 1, len(feedHub.Clients))
	ws.Close()
	waitTimeout(t, feedHub.Sync, 5*time.Second)
	feedHub.Sync.L.Unlock()
	assert.Equal(t, 0, len(feedHub.Clients))

	messageHub.Sync.L.Lock()
	ws1 := test_utils.MakeWSConnection(t, s, "/ws/1")
	waitTimeout(t, messageHub.Sync, 5*time.Second)
	messageHub.Sync.L.Unlock()
	assert.Equal(t, 1, len(messageHub.Clients), "incorrecct number of feeds")
	assert.Equal(t, 1, len(messageHub.Clients["1"]), "incorrecct number of messages")

	messageHub.Sync.L.Lock()
	ws2 := test_utils.MakeWSConnection(t, s, "/ws/1")
	waitTimeout(t, messageHub.Sync, 5*time.Second)
	messageHub.Sync.L.Unlock()
	assert.Equal(t, 1, len(messageHub.Clients), "incorrecct number of feeds")
	assert.Equal(t, 2, len(messageHub.Clients["1"]), "incorrecct number of messages")

	messageHub.Sync.L.Lock()
	ws3 := test_utils.MakeWSConnection(t, s, "/ws/2")
	waitTimeout(t, messageHub.Sync, 5*time.Second)
	messageHub.Sync.L.Unlock()
	assert.Equal(t, 2, len(messageHub.Clients), "incorrecct number of feeds")
	assert.Equal(t, 1, len(messageHub.Clients["2"]), "incorrecct number of messages")

	messageHub.Sync.L.Lock()
	ws1.Close()
	waitTimeout(t, messageHub.Sync, 5*time.Second)
	messageHub.Sync.L.Unlock()
	assert.Equal(t, 2, len(messageHub.Clients), "incorrecct number of feeds")
	assert.Equal(t, 1, len(messageHub.Clients["1"]), "incorrecct number of messages")

	messageHub.Sync.L.Lock()
	ws2.Close()
	waitTimeout(t, messageHub.Sync, 5*time.Second)
	messageHub.Sync.L.Unlock()
	assert.Equal(t, 1, len(messageHub.Clients), "incorrecct number of feeds")

	messageHub.Sync.L.Lock()
	ws3.Close()
	waitTimeout(t, messageHub.Sync, 5*time.Second)
	messageHub.Sync.L.Unlock()
	assert.Equal(t, 0, len(messageHub.Clients), "incorrecct number of feeds")
}
