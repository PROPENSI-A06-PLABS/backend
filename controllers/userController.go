package controllers

import (
	"attendance-payroll-app/initializers"
	"attendance-payroll-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	//get data from body
	var body struct {
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

	c.Bind(&body)

	newUser := models.User{
		Username:         body.Username,
		FullName:         body.FullName,
		Email:            body.Email,
		PhoneNumber:      body.PhoneNumber,
		Password:         body.Password,
		Division:         body.Division,
		Status:           body.Status,
		BaseSalary:       body.BaseSalary,
		ContractDocument: body.ContractDocument,
		CVDocument:       body.CVDocument,
		ProfilePhoto:     body.ProfilePhoto,
	}

	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//return

	c.JSON(http.StatusOK, gin.H{
		"user": newUser,
	})
}
