package models

import (
	"github.com/jinzhu/gorm"
)

// Image model representing images table
type Image struct {
	gorm.Model
	OriginalName string `gorm:"type:varchar(121); not null"`
	Name         string `gorm:"type:varchar(121); not null"`
	Tags         []byte `gorm:"type:json"`
	UserID       uint
	User         User    `gorm:"foreignkey:UserID"`
	BookmarkedBy []*User `gorm:"many2many:bookmarked_images"`
	UploadPath   string  `gorm:"type:mediumtext; not null"`
}

// TableName Set Image's table name to be `images`
func (Image) TableName() string {
	return "images"
}
