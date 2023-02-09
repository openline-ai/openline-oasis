package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/routes/chatHub"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
	"golang.org/x/oauth2"
	"log"
	"strings"
)

// Run will start the server
func Run(conf *c.Config, fh *chatHub.Hub, oauthConfig *oauth2.Config) {
	router := getRouter(conf, fh, oauthConfig)
	if err := router.Run(conf.Service.ServerAddress); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func getRouter(conf *c.Config, fh *chatHub.Hub, oauthConfig *oauth2.Config) *gin.Engine {
	router := gin.New()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = strings.Split(conf.Service.CorsUrl, " ")
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS", "POST", "GET")
	corsConfig.AddAllowHeaders("WebChatApiKey")

	router.Use(cors.New(corsConfig))
	route := router.Group("/api/v1/")

	df := util.MakeDialFactory(conf)
	addMailRoutes(conf, df, route)
	AddWebSocketRoutes(route, fh, conf.WebChat.PingInterval)
	AddWebChatRoutes(conf, df, route)
	route2 := router.Group("/")

	addHealthRoutes(route2)
	return router
}
