package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/openline-ai/openline-grpc"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type CasePostRequest struct {
	Username string
	Message  string
}

func addCaseRoutes(rg *gin.RouterGroup) {
	client := createClient()
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

		// Contact the server and print out its response.
		omsg := &pb.OmniMessage{Type: pb.MessageType_MESSAGE,
			Username:  req.Username,
			Direction: pb.MessageDirection_INBOUND,
			Message:   req.Message,
			Channel:   pb.MessageChannel_WIDGET,
		}
		res, err := client.SaveMessage(c, omsg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res),
		})
	})
}

func createClient() pb.MessageStoreClient {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:9013", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMessageStoreClient(conn)
	return client
}
