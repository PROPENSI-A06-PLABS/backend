package models

import "time"

type Leave struct {
	Id         		uint 		`gorm:"primaryKey"`
	StartDate   	time.Time	
	EndDate    		time.Time		
	Status     		bool		
	Note       		string		
	Type       		string		
	Feedback   		string		
	UserID     		uint 		`gorm:"foreignKey:UserID"`
	ApproverID 		uint 		`gorm:"foreignKey:UserID"`
}
