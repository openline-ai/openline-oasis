package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cr "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	c "github.com/openline-ai/openline-oasis/packages/server/oasis-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/FeedHub"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/routes/MessageHub"
	"github.com/openline-ai/openline-oasis/packages/server/oasis-api/util"
	"log"
	"strings"
)

// Run will start the server
func ConfigureRoutes(conf *c.Config, commonRepositories *cr.Repositories, fh *FeedHub.FeedHub, mh *MessageHub.MessageHub) {
	router := gin.New()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = strings.Split(conf.Service.CorsUrl, " ")
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS", "POST", "GET")

	router.Use(cors.New(corsConfig))

	route := router.Group("/")
	route.Use(service.ApiKeyCheckerHTTP(commonRepositories.AppKeyRepo, service.OASIS_API))
	route.Use(service.UserToTenantEnhancer(commonRepositories.UserRepo))

	df := util.MakeDialFactory(conf)
	addFeedRoutes(route, conf, df)
	addCallCredentialRoutes(route, conf)
	addLoginRoutes(route)

	// TODO: a different typ of auth for websockets
	route2 := router.Group("/")
	AddWebSocketRoutes(route2, fh, mh, conf.WebRTC.PingInterval)

	// no api key (or cors) for health check
	route3 := router.Group("/")
	addHealthRoutes(route3)

	if err := router.Run(conf.Service.ServerAddress); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
