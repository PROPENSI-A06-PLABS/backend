package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username           string
	FullName           string
	Email              string
	PhoneNumber        int
	Password           string
	Division           string
	Status             bool
	BaseSalary         int
	ContractDocument   string
	CVDocument         string
	ProfilePhoto       string
	Role               string
	Identifier         bool
	Attendance         []Attendance   `gorm:"foreignKey:UserID"`
	ApprovedAttendance []Attendance   `gorm:"foreignKey:ApproverID"`
	Leave              []Leave        `gorm:"foreignKey:UserID"`
	ApprovedLeave      []Leave        `gorm:"foreignKey:ApproverID"`
	Payroll            []Payroll      `gorm:"foreignKey:UserID"`
	ApprovedPayroll    []Payroll      `gorm:"foreignKey:ApproverID"`
	Reimburse          []Reimburse    `gorm:"foreignKey:UserID"`
	ApprovedReimburse  []Reimburse    `gorm:"foreignKey:ApproverID"`
	Announcement       []Announcement `gorm:"foreignKey:UserID"`
}
