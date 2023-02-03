package repository

import (
	"encoding/json"
	"github.com/openline-ai/openline-oasis/packages/server/channels-api/repository/entity"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GmailAuthTokensRepository interface {
	Exists(email string) (bool, error)
	Save(gmailAuthToken *entity.GmailAuthToken) (*entity.GmailAuthToken, error)
	Get(email string) (*oauth2.Token, error)
}

type gmailAuthTokensRepository struct {
	db *gorm.DB
}

func NewGmailAuthTokensRepository(db *gorm.DB) GmailAuthTokensRepository {
	return &gmailAuthTokensRepository{
		db: db,
	}
}

func (r *gmailAuthTokensRepository) Save(gmailAuthToken *entity.GmailAuthToken) (*entity.GmailAuthToken, error) {
	result := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"token"}),
	}).Create(&gmailAuthToken)

	if result.Error != nil {
		return nil, result.Error
	}

	return gmailAuthToken, nil
}

func (r *gmailAuthTokensRepository) Exists(email string) (bool, error) {
	result := r.db.Where("email = ?", email).Take(&entity.GmailAuthToken{})

	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

func (r *gmailAuthTokensRepository) Get(email string) (*oauth2.Token, error) {
	element := &entity.GmailAuthToken{}
	result := r.db.Where("email = ?", email).Take(element)

	if result.Error != nil {
		return nil, result.Error
	}

	tok := &oauth2.Token{}
	json.Unmarshal([]byte(element.Token), tok)

	return tok, nil
}