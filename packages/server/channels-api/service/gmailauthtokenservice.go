package service

import (
	"context"
	"encoding/json"
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	proto "github.com/openline-ai/openline-oasis/packages/server/channels-api/proto/generated"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/repository"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/repository/entity"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/routes"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
	"golang.org/x/oauth2"
)

type gmailAuthTokenService struct {
	proto.UnimplementedGmailAuthTokenServiceServer
	conf        *c.Config
	df          util.DialFactory
	repo        *repository.PostgresRepositories
	oauthConfig *oauth2.Config
}

func NewGmailAuthTokenService(c *c.Config, df util.DialFactory, repository *repository.PostgresRepositories, oauthConfig *oauth2.Config) *gmailAuthTokenService {
	gats := new(gmailAuthTokenService)
	gats.conf = c
	gats.repo = repository
	gats.df = df
	gats.oauthConfig = oauthConfig
	return gats
}
func (c *gmailAuthTokenService) GetGmailAuthUrl(ctx context.Context, state *proto.GmailStateInfo) (*proto.GmailAuthUrl, error) {
	gmailState := &routes.GmailState{
		Email:       state.Email,
		RedirectURL: state.RedirectUrl,
	}

	bytes, err := json.Marshal(gmailState)
	if err != nil {
		return nil, err
	}
	authURL := c.oauthConfig.AuthCodeURL(string(bytes), oauth2.AccessTypeOffline)
	return &proto.GmailAuthUrl{Url: authURL}, nil

}

func (c gmailAuthTokenService) SetGmailAuth(ctx context.Context, cred *proto.GmailCredential) (*proto.EventEmpty, error) {
	gmailAuthToken := entity.GmailAuthToken{
		Email: cred.Email,
		Token: cred.Token,
	}
	result := c.repo.GmailAuthTokensRepository.Save(&gmailAuthToken)
	if result.Error != nil {
		return nil, result.Error
	}

	return &proto.EventEmpty{}, nil
}
