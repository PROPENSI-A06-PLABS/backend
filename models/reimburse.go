package models

import "gorm.io/gorm"

type Reimburse struct {
	gorm.Model
	Status     		string 			
	BillStatus 		bool			
	Fee        		int				
	Details    		string			
	Document   		string			
	Feedback   		string			
	UserID     		uint 			`gorm:"foreignKey:UserID"`
	ApproverID 		uint 			`gorm:"foreignKey:UserID"`
}
