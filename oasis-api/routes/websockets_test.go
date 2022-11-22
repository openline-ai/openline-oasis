package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/url"
	"openline-ai/oasis-api/hub"
	"testing"
)

var wsRouter *gin.Engine

func httpToWS(t *testing.T, u string) string {
	t.Helper()

	wsURL, err := url.Parse(u)
	if err != nil {
		t.Fatal(err)
	}

	switch wsURL.Scheme {
	case "http":
		wsURL.Scheme = "ws"
	case "https":
		wsURL.Scheme = "wss"
	}

	return wsURL.String()
}

func makeWSCinnection(t *testing.T, server *httptest.Server, path string) *websocket.Conn {
	wsURL := httpToWS(t, server.URL) + path

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	return ws
}
func newWSServer(t *testing.T) *httptest.Server {

	server := httptest.NewServer(wsRouter)

	return server
}

var feedHub *hub.FeedHub
var messageHub *hub.MessageHub

func setup(t *testing.T) {

	fh := hub.NewFeedHub()
	go fh.RunFeedHub()
	feedHub = fh

	mh := hub.NewMessageHub()
	go mh.RunMessageHub()
	messageHub = mh

	wsRouter = gin.Default()
	route := wsRouter.Group("/")
	addWebSocketRoutes(route, fh, mh)

	t.Cleanup(func() {
		mh.MessageBroadcast <- hub.MessageItem{Id: "quit"}
		_ = <-mh.MessageBroadcast
		fh.FeedBroadcast <- hub.MessageFeed{ContactId: "quit"}
		_ = <-fh.FeedBroadcast

	})
}

func TestWebsocket(t *testing.T) {
	setup(t)
	s := newWSServer(t)
	defer s.Close()
	ws := makeWSCinnection(t, s, "/ws")
	assert.Equal(t, 1, len(feedHub.Clients))
	ws.Close()
	assert.Equal(t, 0, len(feedHub.Clients))
}
