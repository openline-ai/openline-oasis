package routes

import (
	"github.com/caarlos0/env/v6"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	c "openline-ai/oasis-api/config"
)

var router = gin.Default()

// Run will start the server
func Run(addr string) {
	getRoutes()
	if err := router.Run(addr); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func getRoutes() {
	conf := c.Config{}
	env.Parse(&conf)
	v1 := router.Group("/")
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{conf.Service.CorsUrl}
	corsConfig.AllowCredentials = true
	/*	corsConfig.AddAllowHeaders("Origin")
		corsConfig.AddAllowMethods("POST")
		corsConfig.AddAllowMethods("OPTIONS")
		corsConfig.AllowOriginFunc = func(origin string) bool {
			return true
		}
		corsConfig.MaxAge = 12 * time.Hour*/

	router.Use(cors.New(corsConfig))
	addCaseRoutes(v1)

}
