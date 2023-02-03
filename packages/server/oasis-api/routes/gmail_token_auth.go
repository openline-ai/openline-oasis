package routes

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	chProto "github.com/openline-ai/openline-oasis/packages/server/channels-api/proto/generated"
	c "github.com/openline-ai/openline-oasis/packages/server/oasis-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/util"
	"google.golang.org/grpc/metadata"
)

type GmailAuthUrl struct {
	AuthUrl string `json:"auth_url"`
}

type GmailAuthExists struct {
	Exists bool `json:"exists"`
}

func addGmailTokenAuthRoutes(rg *gin.RouterGroup, conf *c.Config, df util.DialFactory) {
	rg.GET("/gmail_token/auth_url", func(c *gin.Context) {
		email := c.Query("email")
		url := c.Query("url")

		// inform the channel api a new message
		channelsConn := util.GetChannelsConnection(c, df)
		defer util.CloseChannelsConnection(channelsConn)
		gmailAuthClient := chProto.NewGmailAuthTokenServiceClient(channelsConn)

		channelsCtx := context.Background()
		channelsCtx = metadata.AppendToOutgoingContext(channelsCtx, service.UsernameHeader, c.GetHeader(service.UsernameHeader))

		authUrl, err := gmailAuthClient.GetGmailAuthUrl(channelsCtx, &chProto.GmailStateInfo{
			Email:       email,
			RedirectUrl: url,
		})

		if err != nil {
			c.JSON(500, gin.H{"msg": fmt.Sprintf("failed to build auth url: %v", err.Error())})
			return
		}
		c.JSON(200, &GmailAuthUrl{AuthUrl: authUrl.Url})
	})

	rg.GET("/gmail_token/exists", func(c *gin.Context) {
		email := c.Query("email")

		// inform the channel api a new message
		channelsConn := util.GetChannelsConnection(c, df)
		defer util.CloseChannelsConnection(channelsConn)
		gmailAuthClient := chProto.NewGmailAuthTokenServiceClient(channelsConn)

		channelsCtx := context.Background()
		channelsCtx = metadata.AppendToOutgoingContext(channelsCtx, service.UsernameHeader, c.GetHeader(service.UsernameHeader))

		exists, err := gmailAuthClient.CheckGmailActive(channelsCtx, &chProto.GmailActiveReq{Email: email})
		if err != nil {
			c.JSON(500, gin.H{"msg": fmt.Sprintf("Failed to query Gmail status: %v", err.Error())})
			return
		}

		c.JSON(200, &GmailAuthExists{Exists: exists.Exists})

	})
}
