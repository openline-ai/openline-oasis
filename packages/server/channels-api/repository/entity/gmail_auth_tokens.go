package entity

type GmailAuthToken struct {
	ID    string `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"id"`
	email string `gorm:"type:varchar(50);NOT NULL" json:"email" binding:"required"`
	token string `gorm:"type:text;NOT NULL" json:"token" binding:"required"`
}
