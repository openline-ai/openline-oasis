package repository

import (
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/repository/helper"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/repository/entity"
	"gorm.io/gorm"
)

type GmailAuthTokensRepository interface {
	Save(gmailAuthToken *entity.GmailAuthToken) helper.QueryResult
}

type gmailAuthTokensRepository struct {
	db *gorm.DB
}

func NewGmailAuthTokensRepository(db *gorm.DB) GmailAuthTokensRepository {
	return &gmailAuthTokensRepository{
		db: db,
	}
}

func (r *gmailAuthTokensRepository) Save(gmailAuthToken *entity.GmailAuthToken) helper.QueryResult {
	result := r.db.Create(&gmailAuthToken)

	if result.Error != nil {
		return helper.QueryResult{Error: result.Error}
	}

	return helper.QueryResult{Result: &gmailAuthToken}
}