package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/routes/chatHub"
)

type Message struct {
	Message string `json:"message"`
}

func AddWebSocketRoutes(rg *gin.RouterGroup, hub *chatHub.Hub, pingInterval int) {

	rg.GET("/ws/:username", func(c *gin.Context) {
		username := c.Param("username")

		if username == "" {
			c.JSON(400, gin.H{"msg": "username missing from path"})
			return
		}
		chatHub.ServeWs(username, hub, c.Writer, c.Request, pingInterval)
	})
}
