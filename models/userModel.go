package models

import "time"

// import "gorm.io/gorm"

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	// gorm.Model       // id, created at, and deleted at
	Username         string
	FullName         string
	Email            string
	PhoneNumber      int
	Password         string
	Division         string
	Status           bool
	BaseSalary       int
	ContractDocument string
	CVDocument       string
	ProfilePhoto     string
}
