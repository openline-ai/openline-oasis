package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// add the middleware function
func apiKeyChecker(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiHeader := c.Request.Header.Get("X-Openline-API-KEY")

		if apiHeader != apiKey {
			log.Printf("Invalid api key %s wanted %s", apiHeader, apiKey)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"reason": "Invalid API Key"})
			return
		}
		c.Next()
	}
}
