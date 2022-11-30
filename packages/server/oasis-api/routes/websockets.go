package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"openline-ai/oasis-api/hub"
)

var msgUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var feedUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Message string `json:"message"`
}

func AddWebSocketRoutes(rg *gin.RouterGroup, fh *hub.FeedHub, mh *hub.MessageHub) {

	rg.GET("/ws", func(c *gin.Context) {
		feedUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := feedUpgrader.Upgrade(c.Writer, c.Request, nil)
		if !errors.Is(err, nil) {
			log.Println(err)
			return // do not add invalid ws to the map
		}
		defer ws.Close()
		addFeedHubConn(fh, ws)
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				removeFeedHubConn(fh, ws)
				return
			}
		}
	})

	rg.GET("/ws/:feedId", func(c *gin.Context) {
		feedId := c.Param("feedId")
		log.Println(feedId)
		msgUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := msgUpgrader.Upgrade(c.Writer, c.Request, nil)
		if !errors.Is(err, nil) {
			log.Println(err)
			return // do not add invalid ws to the map
		}
		defer ws.Close()
		addMessageHubConn(mh, feedId, ws)
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				removeMessageHubConn(mh, feedId, ws)
				return
			}
		}
	})
}

func removeFeedHubConn(fh *hub.FeedHub, ws *websocket.Conn) {
	fh.Sync.L.Lock()
	defer fh.Sync.L.Unlock()
	log.Printf("Cleaning Up Feed Websocket")
	delete(fh.Clients, ws)
	fh.Sync.Signal()
}

func addFeedHubConn(fh *hub.FeedHub, ws *websocket.Conn) {
	fh.Sync.L.Lock()
	fh.Clients[ws] = true
	log.Println("Connected!")
	fh.Sync.Signal()
	fh.Sync.L.Unlock()
}

func removeMessageHubConn(mh *hub.MessageHub, feedId string, ws *websocket.Conn) {
	mh.Sync.L.Lock()
	defer mh.Sync.L.Unlock()

	log.Printf("Cleaning Up Message Websocket")
	delete(mh.Clients[feedId], ws)
	if len(mh.Clients[feedId]) == 0 {
		log.Printf("No more ws for feed %s, deleting feed", feedId)
		delete(mh.Clients, feedId)
	}
	mh.Sync.Signal()
}

func addMessageHubConn(mh *hub.MessageHub, feedId string, ws *websocket.Conn) {
	mh.Sync.L.Lock()
	defer mh.Sync.L.Unlock()
	if _, exists := mh.Clients[feedId]; !exists {
		log.Println("making new feed")
		mh.Clients[feedId] = make(map[*websocket.Conn]bool)
	}
	mh.Clients[feedId][ws] = true
	log.Println("Connected!")
	mh.Sync.Signal()
}
