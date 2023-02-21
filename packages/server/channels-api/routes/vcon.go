package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ms "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/model"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
	"log"
	"net/http"
)

func AddVconRoutes(conf *c.Config, df util.DialFactory, rg *gin.RouterGroup) {
	rg.POST("/vcon", func(c *gin.Context) {
		var req model.VCon
		if err := c.BindJSON(&req); err != nil {
			log.Printf("unable to parse json: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("unable to parse json: %v", err.Error()),
			})
			return
		}

		if conf.VCon.ApiKey != c.GetHeader("X-Openline-VCon-Api-Key") {
			c.JSON(http.StatusForbidden, gin.H{"result": "Invalid API Key"})
			return
		}
		message := &ms.InputMessage{
			Type:                    ms.MessageType_VOICE,
			Subtype:                 ms.MessageSubtype_MESSAGE,
			Content:                 &req.Dialog[0].Body,
			Direction:               ms.MessageDirection_INBOUND,
			InitiatorIdentifier:     &fromAddress,
			ThreadId:                &threadId,
			ParticipantsIdentifiers: toStringArr(append(email.To, email.Cc...)),
		}
}
