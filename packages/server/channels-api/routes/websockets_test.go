package routes

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"openline-ai/channels-api/routes/chatHub"
	"openline-ai/channels-api/test_utils"
	"strconv"
	"testing"
	"time"
)

var hub *chatHub.Hub

func setup(t *testing.T) {

	ch := chatHub.NewHub()
	go ch.Run()
	hub = ch
	test_utils.SetupWebSocketServer(ch, AddWebSocketRoutes)
	t.Cleanup(func() {
		ch.Quit <- true
	})
}

func TestWebsocketCleanup(t *testing.T) {
	setup(t)
	s := test_utils.NewWSServer(t)
	defer s.Close()

	var username = "user@example.org"

	numberOfUsers := 100

	var messages = make([]*websocket.Conn, numberOfUsers)
	for i := 0; i < numberOfUsers; i++ {
		messages[i] = test_utils.MakeWSConnection(t, s, "/ws/"+strconv.Itoa(i)+username)
	}

	assert.Eventually(t, func() bool { return len(hub.Clients) == numberOfUsers }, 2*time.Second, 10*time.Millisecond, "incorrect number of messages")
	for _, message := range messages {
		message.Close()
	}

	assert.Eventually(t, func() bool { return len(hub.Clients) == 0 }, 2*time.Second, 10*time.Millisecond, "Message Hub Clients didn't cleanup: ", len(hub.Clients))

}
