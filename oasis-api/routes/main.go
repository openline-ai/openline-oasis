package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	c "openline-ai/oasis-api/config"
	"openline-ai/oasis-api/hub"
)

// Run will start the server
func Run(conf *c.Config, fh *hub.FeedHub, mh *hub.MessageHub) {
	router := getRouter(conf, fh, mh)
	if err := router.Run(conf.Service.ServerAddress); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func getRouter(config *c.Config, fh *hub.FeedHub, mh *hub.MessageHub) *gin.Engine {
	router := gin.New()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{config.Service.CorsUrl}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS", "POST", "GET")

	router.Use(cors.New(corsConfig))

	route := router.Group("/")
	addFeedRoutes(route, config)
	addCallCredentialRoutes(route, config)
	addLoginRoutes(route)
	addWebSocketRoutes(route, fh, mh)
	return router
}
