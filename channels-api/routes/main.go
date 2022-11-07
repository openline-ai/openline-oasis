package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	c "openline-ai/channels-api/config"
)

// Run will start the server
func Run(conf c.Config) {
	router := getRouter(conf)
	if err := router.Run(conf.Service.ServerAddress); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func getRouter(conf c.Config) *gin.Engine {
	router := gin.New()
	route := router.Group("/")
	addMailRoutes(conf, route)
	return router
}
