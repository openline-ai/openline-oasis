package test_utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/routes/chatHub"
	"net/http/httptest"
	"net/url"
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

func MakeWSConnection(t *testing.T, server *httptest.Server, path string) *websocket.Conn {
	wsURL := httpToWS(t, server.URL) + path

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	return ws
}
func NewWSServer(t *testing.T) *httptest.Server {

	server := httptest.NewServer(wsRouter)

	return server
}

func SetupWebSocketServer(fh *chatHub.Hub, socketRoutes func(*gin.RouterGroup, *chatHub.Hub, int)) {
	wsRouter = gin.Default()
	route := wsRouter.Group("/")
	socketRoutes(route, fh, 30)
}
