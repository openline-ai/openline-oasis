package entity

type GmailAuthToken struct {
	ID    string `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"id"`
	Email string `gorm:"uniqueIndex;type:varchar(50);NOT NULL" json:"email" binding:"required"`
	Token string `gorm:"type:text;NOT NULL" json:"token" binding:"required"`
}
