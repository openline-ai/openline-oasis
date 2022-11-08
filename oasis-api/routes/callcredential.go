package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	c "openline-ai/oasis-api/config"
	"openline-ai/oasis-api/util"
	"time"
)

type CallCredentials struct {
	Username string `json: "username"`
	Password string `json: "password"`
	TTL      int    `json: "ttl"`
}

func addCallCredentialRoutes(rg *gin.RouterGroup, conf c.Config) {

	rg.GET("/call_credentials", func(c *gin.Context) {
		expiresTime := time.Now().Second() + conf.WebRTC.TTL
		timeLimitedUser := fmt.Sprintf("%d:%s", expiresTime, c.Query("username"))
		password := util.GetSignature(timeLimitedUser, conf.WebRTC.AuthSecret)
		c.JSON(http.StatusOK, CallCredentials{Username: timeLimitedUser, Password: password, TTL: conf.WebRTC.TTL})

	})
}
