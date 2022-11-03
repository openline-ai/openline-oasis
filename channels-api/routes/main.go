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
	v1 := router.Group("/api/v1")
	router.Use(cors.Default())
	addMailRoutes(v1)
}
