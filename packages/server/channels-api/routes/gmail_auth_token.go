package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/repository"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/repository/entity"
	"golang.org/x/oauth2"
	"net/http"
)

type gmailAuthTokenRoute struct {
	conf         *c.Config
	oauthConfig  *oauth2.Config
	rg           *gin.RouterGroup
	repositories *repository.PostgresRepositories
}

type GmailState struct {
	Email       string `json:"email"`
	RedirectURL string `json:"redirect_url"`
}

func (gatr *gmailAuthTokenRoute) GetToken(ctx *gin.Context) {
	code := ctx.Query("code")
	state := ctx.Query("state")

	tok, err := gatr.oauthConfig.Exchange(context.TODO(), code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Unable to exchange oauth token: %v", err.Error()),
		})
		return
	}
	bytes, err := json.Marshal(tok)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Unable to json encode token: %v", err.Error()),
		})
		return
	}
	stateInfo := &GmailState{}
	if err := json.Unmarshal([]byte(state), &stateInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Unable to parse json state: %v", err.Error()),
		})
		return
	}

	_, err = gatr.repositories.GmailAuthTokensRepository.Save(&entity.GmailAuthToken{
		Email: stateInfo.Email,
		Token: string(bytes),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Unable to Save token to database: %v", err.Error()),
		})
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, stateInfo.RedirectURL)
}

func addGmailAuthTokenRoutes(conf *c.Config, oauthConfig *oauth2.Config, repositories *repository.PostgresRepositories, rg *gin.RouterGroup) {
	gatr := &gmailAuthTokenRoute{
		conf:         conf,
		oauthConfig:  oauthConfig,
		rg:           rg,
		repositories: repositories,
	}
	rg.GET("auth", gatr.GetToken)

}
