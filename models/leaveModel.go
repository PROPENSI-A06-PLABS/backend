package models

import "time"

type Leave struct {
	Id       	uint `gorm:"primaryKey"`
	StarDate 	time.Time
	EndDate 	time.Time
	Status 		bool
	Note 		string
	Type 		string
	Feedback 	string
}