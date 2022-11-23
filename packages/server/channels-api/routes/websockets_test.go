package routes

import (
	"github.com/stretchr/testify/assert"
	"openline-ai/channels-api/hub"
	"openline-ai/channels-api/test_utils"
	"testing"
	"time"
)

var webchatMessageHub *hub.WebChatMessageHub

func setup(t *testing.T) {

	fh := hub.NewWebChatMessageHub()
	go fh.RunWebChatMessageHub()
	webchatMessageHub = fh

	test_utils.SetupWebSocketServer(fh, AddWebSocketRoutes)

	t.Cleanup(func() {
		fh.MessageBroadcast <- hub.WebChatMessageItem{Kill: true}
		_ = <-fh.MessageBroadcast

	})
}

func TestWebsocketCleanup(t *testing.T) {
	setup(t)
	s := test_utils.NewWSServer(t)
	defer s.Close()

	var username1 = "gabi@example.org"
	var username2 = "torrey@example.org"

	ws1 := test_utils.MakeWSConnection(t, s, "/ws/"+username1)
	time.Sleep(500 * time.Millisecond)
	assert.Equal(t, 1, len(webchatMessageHub.Clients), "incorrect number of users")
	assert.Equal(t, 1, len(webchatMessageHub.Clients[username1]), "incorrect number of connections")
	ws2 := test_utils.MakeWSConnection(t, s, "/ws/"+username1)
	assert.Equal(t, 1, len(webchatMessageHub.Clients), "incorrect number of users")
	assert.Equal(t, 2, len(webchatMessageHub.Clients[username1]), "incorrect number of connections")
	ws3 := test_utils.MakeWSConnection(t, s, "/ws/"+username2)
	assert.Equal(t, 2, len(webchatMessageHub.Clients), "incorrect number of users")
	assert.Equal(t, 1, len(webchatMessageHub.Clients[username2]), "incorrect number of connections")

	ws1.Close()
	time.Sleep(2 * time.Second)
	assert.Equal(t, 2, len(webchatMessageHub.Clients), "incorrect number of users")
	assert.Equal(t, 1, len(webchatMessageHub.Clients[username1]), "incorrect number of connections")

	ws2.Close()
	time.Sleep(2 * time.Second)
	assert.Equal(t, 1, len(webchatMessageHub.Clients), "incorrect number of users")

	ws3.Close()
	time.Sleep(2 * time.Second)
	assert.Equal(t, 0, len(webchatMessageHub.Clients), "incorrect number of users")
}
