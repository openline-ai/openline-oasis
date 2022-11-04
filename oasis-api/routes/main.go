package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	c "openline-ai/oasis-api/config"
)

var router = gin.Default()

// Run will start the server
func Run(addr string, config c.Config) {
	router := getRouter(config)
	if err := router.Run(addr); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func getRouter(config c.Config) *gin.Engine {
	router := gin.New()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{config.Service.CorsUrl}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS", "POST")

	router.Use(cors.New(corsConfig))

	v1 := router.Group("/")
	addFeedRoutes(v1, config)
	addLoginRoutes(v1)
	return router
}
