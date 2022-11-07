package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginPostRequest struct {
	Username string `json:"username" binding:"required""`
	Password string `json:"password" binding:"required"`
}

func addLoginRoutes(rg *gin.RouterGroup) {

	rg.POST("/login", func(c *gin.Context) {
		fmt.Println(c.Request.Body)
		var req LoginPostRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		c.JSON(http.StatusOK, req.Username)

	})

	rg.POST("/logout", func(c *gin.Context) {
		var req LoginPostRequest

		c.JSON(http.StatusOK, req.Username)

	})
	rg.POST("/account", func(c *gin.Context) {
		var req LoginPostRequest

		c.JSON(http.StatusOK, req.Username)

	})
}
