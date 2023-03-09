package models

import "gorm.io/gorm"

type Payroll struct {
	gorm.Model
	Status          bool		
	Bonus           int			
	Allowance       int			
	AKS             int			
	ATK             int			
	PPH_21          int			
	BPJS_AKS        int			
	BPJS_ATK        int			
	Loan            int			
	TotalEarnings   int			
	TotalDeductions int			
	TotalHomePay    int			
	UserID          uint 		`gorm:"foreignKey:UserID"`
	ApproverID      uint 		`gorm:"foreignKey:UserID"`
}
