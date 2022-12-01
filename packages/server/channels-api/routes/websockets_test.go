package routes

import (
	"github.com/stretchr/testify/assert"
	"openline-ai/channels-api/hub"
	"openline-ai/channels-api/test_utils"
	"sync"
	"testing"
	"time"
)

var webchatMessageHub *hub.WebChatMessageHub

func setup(t *testing.T) {

	fh := hub.NewWebChatMessageHub()
	go fh.RunWebChatMessageHub(60)
	webchatMessageHub = fh

	test_utils.SetupWebSocketServer(fh, AddWebSocketRoutes)

	t.Cleanup(func() {
		fh.MessageBroadcast <- hub.WebChatMessageItem{Kill: true}
		_ = <-fh.MessageBroadcast

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

	var username1 = "gabi@example.org"
	var username2 = "torrey@example.org"

	webchatMessageHub.Sync.L.Lock()
	ws1 := test_utils.MakeWSConnection(t, s, "/ws/"+username1)
	waitTimeout(t, webchatMessageHub.Sync, 5*time.Second)
	webchatMessageHub.Sync.L.Unlock()
	assert.Equal(t, 1, len(webchatMessageHub.Clients), "incorrect number of users")
	assert.Equal(t, 1, len(webchatMessageHub.Clients[username1]), "incorrect number of connections")

	webchatMessageHub.Sync.L.Lock()
	ws2 := test_utils.MakeWSConnection(t, s, "/ws/"+username1)
	waitTimeout(t, webchatMessageHub.Sync, 5*time.Second)
	webchatMessageHub.Sync.L.Unlock()
	assert.Equal(t, 1, len(webchatMessageHub.Clients), "incorrect number of users")
	assert.Equal(t, 2, len(webchatMessageHub.Clients[username1]), "incorrect number of connections")

	webchatMessageHub.Sync.L.Lock()
	ws3 := test_utils.MakeWSConnection(t, s, "/ws/"+username2)
	waitTimeout(t, webchatMessageHub.Sync, 5*time.Second)
	webchatMessageHub.Sync.L.Unlock()
	assert.Equal(t, 2, len(webchatMessageHub.Clients), "incorrect number of users")
	assert.Equal(t, 1, len(webchatMessageHub.Clients[username2]), "incorrect number of connections")

	webchatMessageHub.Sync.L.Lock()
	ws1.Close()
	waitTimeout(t, webchatMessageHub.Sync, 5*time.Second)
	webchatMessageHub.Sync.L.Unlock()
	assert.Equal(t, 2, len(webchatMessageHub.Clients), "incorrect number of users")
	assert.Equal(t, 1, len(webchatMessageHub.Clients[username1]), "incorrect number of connections")

	webchatMessageHub.Sync.L.Lock()
	ws2.Close()
	waitTimeout(t, webchatMessageHub.Sync, 5*time.Second)
	webchatMessageHub.Sync.L.Unlock()
	assert.Equal(t, 1, len(webchatMessageHub.Clients), "incorrect number of users")

	webchatMessageHub.Sync.L.Lock()
	ws3.Close()
	waitTimeout(t, webchatMessageHub.Sync, 5*time.Second)
	webchatMessageHub.Sync.L.Unlock()
	assert.Equal(t, 0, len(webchatMessageHub.Clients), "incorrect number of users")
}
