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

func AddWebSocketRoutes(rg *gin.RouterGroup, fh *hub.WebChatMessageHub) {

	rg.GET("/ws/:username", func(c *gin.Context) {
		username := c.Param("username")

		if username == "" {
			c.JSON(400, gin.H{"msg": "username missing from path"})
			return
		}
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if !errors.Is(err, nil) {
			log.Println(err)
		}
		defer ws.Close()

		if _, exists := fh.Clients[username]; !exists {
			log.Printf("making new feed for %s", username)
			fh.Clients[username] = make(map[*websocket.Conn]bool)
		}
		fh.Clients[username][ws] = true

		log.Println("Connected!")
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				log.Printf("Cleaning Up Webchat Websocket")
				delete(fh.Clients[username], ws)
				if len(fh.Clients[username]) == 0 {
					log.Printf("No more ws for user %s, deleting feed", username)
					delete(fh.Clients, username)
				}
				return
			}
		}
	})
}
