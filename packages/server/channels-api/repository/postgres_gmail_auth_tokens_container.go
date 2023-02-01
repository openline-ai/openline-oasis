package repository

import "gorm.io/gorm"

type GmailAuthTokensRepository interface {
}

type gmailAuthTokensRepository struct {
	db *gorm.DB
}

func NewGmailAuthTokensRepository(db *gorm.DB) GmailAuthTokensRepository {
	return &gmailAuthTokensRepository{
		db: db,
	}
}
