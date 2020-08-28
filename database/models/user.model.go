package models

import (
	"github.com/jinzhu/gorm"
)

// User models representing users table
type User struct {
	gorm.Model
	Email            string       `gorm:"type:varchar(121); unique; not null"`
	Username         string       `gorm:"type:varchar(121); unique; not null"`
	Password         string       `gorm:"type:varchar(121); not null"`
	Tokens           []*UserToken `gorm:"foreignkey:UserID"`
	Images           []*Image     `gorm:"foreignkey:UserID"`
	BookmarkedImages []*Image     `gorm:"many2many:bookmarked_images"`
}

// TableName Set User's table name to be `users`
func (User) TableName() string {
	return "users"
}
