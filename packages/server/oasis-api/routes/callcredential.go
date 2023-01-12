package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	c "github.com/openline-ai/openline-oasis/packages/server/oasis-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/util"
	"net/http"
	"time"
)

func addCallCredentialRoutes(rg *gin.RouterGroup, conf *c.Config) {

	rg.GET("/call_credentials", func(c *gin.Context) {
		expiresTime := time.Now().Unix() + int64(conf.WebRTC.TTL)
		timeLimitedUser := fmt.Sprintf("%d:%s", expiresTime, c.Query("username"))
		password := util.GetSignature(timeLimitedUser, conf.WebRTC.AuthSecret)
		c.JSON(http.StatusOK, gin.H{"username": timeLimitedUser, "password": password, "ttl": conf.WebRTC.TTL})

	})
}
