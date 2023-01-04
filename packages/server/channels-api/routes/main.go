package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	c "openline-ai/channels-api/config"
	"openline-ai/channels-api/routes/chatHub"
	"openline-ai/channels-api/util"
)

// Run will start the server
func Run(conf *c.Config, fh *chatHub.Hub) {
	router := getRouter(conf, fh)
	if err := router.Run(conf.Service.ServerAddress); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func getRouter(conf *c.Config, fh *chatHub.Hub) *gin.Engine {
	router := gin.New()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
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
