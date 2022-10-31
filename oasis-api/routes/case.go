package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CasePostRequest struct {
	Username string
	Message  string
}

func addCaseRoutes(rg *gin.RouterGroup) {
	caseRoute := rg.Group("/case")
	caseRoute.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "case get")
	})
	caseRoute.POST("/", func(c *gin.Context) {
		var req CasePostRequest
		if err := c.BindJSON(&req); err != nil {
			// DO SOMETHING WITH THE ERROR
		}
		c.JSON(http.StatusOK, "Case POST endpoint. req sent: username "+req.Username+"; Message: "+req.Message)

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(""),
		})
	})
}
