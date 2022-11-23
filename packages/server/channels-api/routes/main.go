package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	c "openline-ai/channels-api/config"
	"openline-ai/channels-api/hub"
	"openline-ai/channels-api/util"
	"strings"
)

// Run will start the server
func Run(conf *c.Config, fh *hub.WebChatMessageHub) {
	router := getRouter(conf, fh)
	if err := router.Run(conf.Service.ServerAddress); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func getRouter(conf *c.Config, fh *hub.WebChatMessageHub) *gin.Engine {
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
	addWebSocketRoutes(route, fh)
	AddWebChatRoutes(conf, route)
	return router
}
