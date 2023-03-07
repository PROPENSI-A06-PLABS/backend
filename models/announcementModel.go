package models

import "gorm.io/gorm"

type Announcement struct {
	gorm.Model
	Division string
	Text     string
	User     User
	UserID   uint `gorm:"foreignKey:UserID"`
}
