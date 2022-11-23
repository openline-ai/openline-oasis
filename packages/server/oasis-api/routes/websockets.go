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
		}
		defer ws.Close()
		fh.Clients[ws] = true
		log.Println("Connected!")
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				log.Printf("Cleaning Up Feed Websocket")
				delete(fh.Clients, ws)
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
		}
		defer ws.Close()
		if _, exists := mh.Clients[feedId]; !exists {
			log.Println("making new feed")
			mh.Clients[feedId] = make(map[*websocket.Conn]bool)
		}
		mh.Clients[feedId][ws] = true
		log.Println("Connected!")
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				log.Printf("Cleaning Up Message Websocket")
				delete(mh.Clients[feedId], ws)
				if len(mh.Clients[feedId]) == 0 {
					log.Printf("No more ws for feed %s, deleting feed", feedId)
					delete(mh.Clients, feedId)
				}
				return
			}
		}
	})
}