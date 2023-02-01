package service

import (
	c "github.com/openline-ai/openline-oasis/packages/server/channels-api/config"
	proto "github.com/openline-ai/openline-oasis/packages/server/channels-api/proto/generated"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/repository"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/util"
)

type gmailAuthTokenService struct {
	proto.UnimplementedGmailAuthTokenServiceServer
	conf *c.Config
	df   util.DialFactory
	repo *repository.PostgresRepositories
}

func NewGmailAuthTokenService(c *c.Config, df util.DialFactory, repository *repository.PostgresRepositories) *gmailAuthTokenService {
	gats := new(gmailAuthTokenService)
	gats.conf = c
	gats.repo = repository
	gats.df = df
	return gats
}
