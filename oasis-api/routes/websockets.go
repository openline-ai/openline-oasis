package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"openline-ai/oasis-api/hub"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Message string `json:"message"`
}

func addWebSocketRoutes(rg *gin.RouterGroup, fh *hub.FeedHub, mh *hub.MessageHub) {

	rg.GET("/ws", func(c *gin.Context) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if !errors.Is(err, nil) {
			log.Println(err)
		}
		defer ws.Close()
		fh.Clients[ws] = true
		read(ws)
		log.Println("Connected!")
	})

	rg.GET("/ws/:feedId", func(c *gin.Context) {
		feedId := c.Param("feedId")
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if !errors.Is(err, nil) {
			log.Println(err)
		}
		defer ws.Close()
		mh.Clients[feedId] = ws
		log.Println("Connected!")
		read(ws)
	})
}

func read(ws *websocket.Conn) {
	for {
		var message Message
		err := ws.ReadJSON(&message)
		if !errors.Is(err, nil) {
			log.Printf("error occurred: %v", err)
			//delete(hub.clients, ws)
			break
		}
		log.Println(message)
	}
}
