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

func addWebSocketRoutes(rg *gin.RouterGroup, fh *hub.FeedHub, mh *hub.MessageHub) {

	rg.GET("/ws", func(c *gin.Context) {
		feedUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := feedUpgrader.Upgrade(c.Writer, c.Request, nil)
		if !errors.Is(err, nil) {
			log.Println(err)
		}
		defer ws.Close()
		fh.Clients[ws] = true
		log.Println("Connected!")
		select {}
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
		mh.Clients[feedId] = ws
		log.Println("Connected!")
		select {}
	})
}
