package services

import (
	"attendance-payroll-app/initializers"
	"attendance-payroll-app/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/go-playground/validator/v10"
)

func CheckIn(c *gin.Context) (models.Attendance, string, int) {

	// get data from body
	newAttendance := models.Attendance{}
	err := c.Bind(&newAttendance)
	if err != nil {
		return newAttendance, "Failed to read body", http.StatusBadRequest
	}

	checkinTime := time.Now()
	var isOntime bool

	if newAttendance.CheckinTime.Hour() < 9 {
		isOntime = true
	} else {
		isOntime = false
	}

	newAttendance.CheckinTime = checkinTime
	newAttendance.Date = checkinTime
	newAttendance.Status = isOntime

	// validate input
	validate := validator.New()
	err = validate.Struct(newAttendance)
	if err != nil {
		return newAttendance, err.Error(), http.StatusBadRequest
	}

	result := initializers.DB.Create(&newAttendance)
	if result.Error != nil {
		return newAttendance, "Failed to create new attendance", http.StatusBadRequest
	}

	return newAttendance, "", http.StatusOK
}

func CheckOut(c *gin.Context) (models.Attendance, error, int) {
	var attendance models.Attendance
	var user models.User

	// get data from body
	id := c.Param("id")
	initializers.DB.First(&user, id)
	
	// retrieve current attendance
	err := initializers.DB.Last(&attendance, "user_id=?", id).Error
	if err != nil {
		return attendance, err, http.StatusBadRequest
	}

	checkoutTime := time.Now()
	attendance.CheckoutTime = checkoutTime

	// update checkout time attribute
	err = initializers.DB.Model(&attendance).Updates(
		models.Attendance{
			CheckoutTime: checkoutTime,
		},
	).Error
	if err != nil {
		return attendance, err, http.StatusBadRequest
	}

	return attendance, nil, http.StatusOK
}

func UpdateAttendance(c *gin.Context) (models.Attendance, error, int) {
	id := c.Param("id")
	attendance := models.Attendance{}
	err := initializers.DB.First(&attendance, id).Error
	if (err != nil){
		return attendance, err, http.StatusBadRequest
	}

	// get user's input
	input := models.Attendance{}
	err = c.Bind(&input)
	if (err != nil){
		return attendance, err, http.StatusBadRequest
	}

	// update attendance
	err = initializers.DB.Model(&attendance).Updates(
		models.Attendance{
			CheckinTime: input.CheckinTime,
			CheckoutTime: input.CheckoutTime,
			Date: input.Date,
			Approval: input.Approval,
			Status: input.Status,
			ApproverID: input.ApproverID,
		},
	).Error
	if (err != nil){
		return attendance, err, http.StatusBadRequest
	}
	input.Id = attendance.Id
	return input, nil, http.StatusOK
}

func GetAllAttendance(c *gin.Context) ([]models.Attendance, error, int) {
	// get all attendance
	attendances := []models.Attendance{}
	err := initializers.DB.Find(&attendances).Error
	if err != nil {
		return attendances, err, http.StatusBadRequest
	}

	return attendances, nil, http.StatusOK
}

func GetUserAttendance(c *gin.Context) ([]models.Attendance, error, int) {
	id := c.Param("id")
	attendances := []models.Attendance{}
	err := initializers.DB.Find(&attendances, "user_id=?", id).Error
	if err != nil {
		return attendances, err, http.StatusBadRequest
	}

	return attendances, nil, http.StatusOK
}

func DeleteAttendance(c *gin.Context) (error, int) {
	// delete user by id
	id := c.Param("id")
	attendance := models.Attendance{}
	err := initializers.DB.Delete(&attendance, id).Error
	if err != nil {
		return err, http.StatusBadRequest
	}
	return nil, http.StatusOK
}