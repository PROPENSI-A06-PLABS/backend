package services

import (
	"attendance-payroll-app/initializers"
	"attendance-payroll-app/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

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