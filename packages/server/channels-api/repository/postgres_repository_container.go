package repository

import (
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/repository/entity"
	"gorm.io/gorm"
	"log"
)

type PostgresRepositories struct {
	GmailAuthTokensRepository GmailAuthTokensRepository
}

func InitRepositories(db *gorm.DB) *PostgresRepositories {
	p := &PostgresRepositories{
		GmailAuthTokensRepository: NewGmailAuthTokensRepository(db),
	}

	err := db.AutoMigrate(&entity.GmailAuthToken{})
	if err != nil {
		log.Print(err)
		panic(err)
	}

	return p
}
