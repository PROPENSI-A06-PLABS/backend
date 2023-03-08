package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	Id             uint `gorm:"primaryKey"` //blm implement id plabs
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Username       string
	Name           string `gorm:"not null"`
	BornDay        time.Time
	IdentityNumber int
	AddressDetail  string
	Email          string `gorm:"not null"`
	PLABSMail      string
	Github         string
	Gitlab         string
	//bank acc belom
	Phone    int `gorm:"not null"`
	Password string
	// Division string
	Position string `gorm:"not null"`
	Status   bool   `gorm:"not null"`
	// BaseSalary       int
	Role               string
	Identifier         bool
	StartWork          time.Time `gorm:"not null"`
	Tenure             string    `gorm:"not null"`
	ContractType       string    `gorm:"not null"`
	GrossSalary        int       `gorm:"not null"`
	NPWP               string
	JKN_KIS            int
	KPJ                int
	KTPDocument        string
	NPWPDocument       string
	CVDocument         string
	ContractDocument   string
	ProfilePhoto       string
	ExtraInfo          string
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
