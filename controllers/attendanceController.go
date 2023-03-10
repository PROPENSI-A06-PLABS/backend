package controllers

import (
	"attendance-payroll-app/initializers"
	"attendance-payroll-app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckIn(c *gin.Context) {
	// get data from body
	var body struct {
		CheckinTime		time.Time
		Date 					time.Time
		UserID       	uint
		ApproverID    uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})

		return
	}

	currentDate := time.Now()

	// create attendance
	newAttendance := models.Attendance{
		CheckinTime: 		currentDate,
		Date: 					currentDate,
		UserID: 				body.UserID,	
		ApproverID: 		body.ApproverID,	
	}

	result := initializers.DB.Create(&newAttendance)

	var user models.User

	initializers.DB.First(&user, newAttendance.UserID)
	
	initializers.DB.Model(&user).Association("Attendance").Append(&newAttendance)
	initializers.DB.Preload("Attendance").Find(&user, newAttendance.UserID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create attendance",
		})
	}

	//return
	json, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Entering error block")
	}
	fmt.Println(string(json))

	c.JSON(http.StatusOK, gin.H{
		"user": newAttendance,
	})
}

func GetAllAttendance(c *gin.Context) {
	attendance := []models.Attendance{}
	initializers.DB.Find(&attendance)
	c.JSON(http.StatusOK, attendance)
}

func GetUserAttendance(c *gin.Context) {
	id := c.Param("id")
	attendance := []models.Attendance{}
	initializers.DB.Find(&attendance, "user_id=?", id)
	c.JSON(http.StatusOK, attendance)
}

// func CheckOut(c *gin.Context) {

// }