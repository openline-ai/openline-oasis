package routes

//import (
//	"github.com/gorilla/websocket"
//	"github.com/stretchr/testify/assert"
//	"openline-ai/oasis-api/routes/FeedHub"
//	"openline-ai/oasis-api/routes/MessageHub"
//	"openline-ai/oasis-api/test_utils"
//	"strconv"
//	"testing"
//	"time"
//)
//
//var feedHub *FeedHub.FeedHub
//var messageHub *MessageHub.MessageHub
//
//func setup(t *testing.T) {
//
//	fh := FeedHub.NewFeedHub()
//	go fh.Run()
//	feedHub = fh
//
//	mh := MessageHub.NewMessageHub()
//	go mh.Run()
//	messageHub = mh
//
//	test_utils.SetupWebSocketServer(fh, mh, AddWebSocketRoutes)
//
//	t.Cleanup(func() {
//		mh.Quit <- true
//		fh.Quit <- true
//	})
//}
//
//func TestWebsocketCleanup(t *testing.T) {
//	setup(t)
//	s := test_utils.NewWSServer(t)
//	defer s.Close()
//
//	numberOfFeeds := 20
//
//	var feeds = make([]*websocket.Conn, numberOfFeeds)
//	for i := 0; i < numberOfFeeds; i++ {
//		feeds[i] = test_utils.MakeWSConnection(t, s, "/ws")
//	}
//
//	assert.Eventually(t, func() bool { return len(feedHub.Clients) == numberOfFeeds }, 2*time.Second, 10*time.Millisecond)
//
//	for _, feed := range feeds {
//		feed.Close()
//	}
//
//	assert.Eventually(t, func() bool { return len(feedHub.Clients) == 0 }, 2*time.Second, 10*time.Millisecond, "Feed Hub clients didn't cleanup")
//
//	numberOfMessages := 100
//
//	var messages = make([]*websocket.Conn, numberOfMessages)
//	for i := 0; i < numberOfMessages; i++ {
//		messages[i] = test_utils.MakeWSConnection(t, s, "/ws/"+strconv.Itoa(i))
//	}
//
//	assert.Eventually(t, func() bool { return len(messageHub.Clients) == numberOfMessages }, 2*time.Second, 10*time.Millisecond, "incorrect number of messages")
//	for _, message := range messages {
//		message.Close()
//	}
//
//	assert.Eventually(t, func() bool { return len(messageHub.Clients) == 0 }, 2*time.Second, 10*time.Millisecond, "Message Hub Clients didn't cleanup")
//}
