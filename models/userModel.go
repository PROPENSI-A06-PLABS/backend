package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	Id        			uint 			`gorm:"primaryKey"` //blm implement id plabs
	CreatedAt 			time.Time  		
	UpdatedAt 			time.Time  		
	DeletedAt 			gorm.DeletedAt 	`gorm:"index"`
	
	// section 1
	Name           		string 			`gorm:"not null" json:"Name" validate:"required"`
	Phone    			string    		`gorm:"not null" json:"Phone" validate:"required"`
	Email          		string 			`gorm:"not null" json:"Email" validate:"required,email"` 
	BornDay        		time.Time 		
	AddressDetail  		string			
	IdentityNumber 		int		

	// section 2
	AccountNumber 		int 			// no rek
	NPWP               	string
	KPJ                	int				
	Github         		string			
	BankName 			string
	JKN_KIS            	int				
	Gitlab         		string		
	ExtraInfo          	string			

	// section 3
	PLABSMail      		string			
	Position 			string 			`gorm:"not null" json:"Position" validate:"required"`
	GrossSalary       	int       		`gorm:"not null" json:"GrossSalary" validate:"required"`
	Role               	string			`gorm:"not null" json:"Role" validate:"required,eq=ADMIN|eq=MANAGER|eq=HR|eq=STAFF"`
	ContractType       	string    		`gorm:"not null" json:"ContractType" validate:"required"`
	Password 			string 			`gorm:"not null" json:"Password" validate:"required,min=7"`
	StartWork          	time.Time 		`gorm:"not null" json:"StartWork" validate:"required"`
	Status   			bool   			`gorm:"not null"`  // active - inactive

	// document
	NPWPDocument       string			
	KTPDocument        string			
	CVDocument         string			
	ContractDocument   string			
	ProfilePhoto       string			

	// additional
	Identifier         bool				
	Tenure             string    		`gorm:"not null"`
	Attendance         []Attendance   	`gorm:"foreignKey:UserID"`
	ApprovedAttendance []Attendance   	`gorm:"foreignKey:ApproverID"`
	Leave              []Leave        	`gorm:"foreignKey:UserID"`
	ApprovedLeave      []Leave        	`gorm:"foreignKey:ApproverID"`
	Payroll            []Payroll      	`gorm:"foreignKey:UserID"`
	ApprovedPayroll    []Payroll      	`gorm:"foreignKey:ApproverID"`
	Reimburse          []Reimburse    	`gorm:"foreignKey:UserID"`
	ApprovedReimburse  []Reimburse    	`gorm:"foreignKey:ApproverID"`
	Announcement       []Announcement 	`gorm:"foreignKey:UserID"`
}
