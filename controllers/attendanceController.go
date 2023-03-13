package controllers

import (
	"attendance-payroll-app/initializers"
	"attendance-payroll-app/models"
	"attendance-payroll-app/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckIn(c *gin.Context) {
	// get data from body
	var body struct {
		// CheckinTime time.Time
		// Date        time.Time
		UserID uint
		// ApproverID  *uint
		Location string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})

		return
	}

	checkinTime := time.Now()

	// create attendance
	newAttendance := models.Attendance{
		CheckinTime: checkinTime,
		Date:        checkinTime,
		UserID:      body.UserID,
		ApproverID:  nil,
		Location:    body.Location,
	}

	result := initializers.DB.Create(&newAttendance)

	var user models.User

	initializers.DB.First(&user, newAttendance.UserID)
	initializers.DB.Model(&user).Association("Attendance").Append(&newAttendance)
	// initializers.DB.Preload("Attendance").Find(&user, newAttendance.UserID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create attendance",
		})
	}

	// json, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println("Entering error block")
	// }
	// fmt.Println(string(json))

	c.JSON(http.StatusOK, gin.H{
		"user": user,
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

func CheckOut(c *gin.Context) {
	var attendance models.Attendance
	var user models.User

	id := c.Param("id")
	initializers.DB.First(&user, id)
	initializers.DB.Last(&attendance, "user_id=?", id)

	checkoutTime := time.Now()

	// update attendance
	initializers.DB.Model(&attendance).Updates(
		models.Attendance{
			CheckoutTime: checkoutTime,
		},
	)

	initializers.DB.Model(&user).Association("Attendance").Append(&attendance)

	// json, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println("Entering error block")
	// }
	// fmt.Println(string(json))

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UpdateAttendance(c *gin.Context) {
	attendance, err, status := services.UpdateAttendance(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(status, attendance)
	}
}
