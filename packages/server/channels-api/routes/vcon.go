package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/model"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
	"log"
	"net/http"
)

func AddVconRoutes(conf *c.Config, df util.DialFactory, rg *gin.RouterGroup) {
	rg.POST("/vcon", func(c *gin.Context) {
		var req model.VCon
		if err := c.BindJSON(&req); err != nil {
			log.Printf("unable to parse json: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to parse json: %v", err.Error()),
			})
			return
		}

		if conf.WebChat.ApiKey != c.GetHeader("WebChatApiKey") {
			c.JSON(http.StatusForbidden, gin.H{"result": "Invalid API Key"})
			return
		}
}
