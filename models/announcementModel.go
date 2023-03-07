package models

import "gorm.io/gorm"

type Announcement struct {
	gorm.Model
	Division 	string 
	Text 		string
	User User
}