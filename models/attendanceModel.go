package models

import (
	"time"
)

type Attendance struct {
	Id           	uint 			`gorm:"primaryKey"`
	CheckinTime  	time.Time
	CheckoutTime 	time.Time
	Date         	time.Time
	Approval 			bool
	Status       	bool
	UserID       	uint 			`gorm:"foreignKey:UserID"`
	ApproverID   	*uint 			`gorm:"foreignKey:UserID"`
}
