package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"openline-ai/channels-api/hub"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Message string `json:"message"`
}

func addWebSocketRoutes(rg *gin.RouterGroup, fh *hub.WebChatMessageHub) {

	rg.GET("/ws/:username", func(c *gin.Context) {
		var username string
		if err := c.ShouldBindUri(&username); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if !errors.Is(err, nil) {
			log.Println(err)
		}
		defer ws.Close()

		fh.Clients[username] = ws

		select {}

	})
}
