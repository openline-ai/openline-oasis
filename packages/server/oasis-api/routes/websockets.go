package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	FeedHub "github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/FeedHub"
	MessageHub "github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/MessageHub"
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

func AddWebSocketRoutes(rg *gin.RouterGroup, fh *FeedHub.FeedHub, mh *MessageHub.MessageHub, pingInterval int) {

	rg.GET("/ws", func(c *gin.Context) {
		FeedHub.ServeFeedWs(fh, c.Writer, c.Request, pingInterval)
	})

	rg.GET("/ws/:feedId", func(c *gin.Context) {
		feedId := c.Param("feedId")
		if feedId == "" {
			c.JSON(400, gin.H{"msg": "feedId missing from path"})
			return
		}
		MessageHub.ServeMessageWs(feedId, mh, c.Writer, c.Request, pingInterval)
	})
}
