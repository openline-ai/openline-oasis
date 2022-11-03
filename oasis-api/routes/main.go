package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
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
	v1 := router.Group("/")
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3006"}
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
