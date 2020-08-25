package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserToken model representing users_tokens table
type UserToken struct {
	gorm.Model
	UserID       uint
	User         User      `gorm:"foreignkey:UserID"`
	AccessToken  string    `gorm:"type:longtext; not null"`
	RefreshToken string    `gorm:"type:varchar(121); not null"`
	ExpiresAt    time.Time `gorm:"type:datetime; not null"`
}

// TableName Set UserToken's table name to be `users_tokens`
func (UserToken) TableName() string {
	return "users_tokens"
}
