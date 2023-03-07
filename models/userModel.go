package models

import "gorm.io/gorm"


type User struct {
	gorm.Model       
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
	Role 	    	 string
	Identifier 		 bool
	Attendance 		[]Attendance
}
