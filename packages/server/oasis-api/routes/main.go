package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	c "github.com/openline-ai/openline-oasis/packages/server/oasis-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/FeedHub"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/MessageHub"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/util"
	"log"
	"strings"
)

// Run will start the server
func Run(conf *c.Config, fh *FeedHub.FeedHub, mh *MessageHub.MessageHub) {
	router := getRouter(conf, fh, mh)
	if err := router.Run(conf.Service.ServerAddress); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func getRouter(config *c.Config, fh *FeedHub.FeedHub, mh *MessageHub.MessageHub) *gin.Engine {
	router := gin.New()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = strings.Split(config.Service.CorsUrl, " ")
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS", "POST", "GET")

	router.Use(cors.New(corsConfig))

	route := router.Group("/")
	route.Use(apiKeyChecker(config.Service.ApiKey))

	df := util.MakeDialFactory(config)
	addFeedRoutes(route, config, df)
	addCallCredentialRoutes(route, config)
	addLoginRoutes(route)

	// TODO: a different typ of auth for websockets
	route2 := router.Group("/")
	AddWebSocketRoutes(route2, fh, mh, config.WebRTC.PingInterval)

	// no api key (or cors) for health check
	route3 := router.Group("/")
	addHealthRoutes(route3)
	return router
}
