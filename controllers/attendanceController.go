package controllers

import (
	"attendance-payroll-app/services"
	"github.com/gin-gonic/gin"
)

func CheckIn(c *gin.Context) {
	attendance, err, status := services.CheckIn(c)
	if err != "" {
		c.JSON(status, gin.H{
			"message": err,
		})
	} else {
		c.JSON(status, attendance)
	}
}

func CheckOut(c *gin.Context) {
	attendance, err, status := services.CheckOut(c)
	if err !=  nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(status, attendance)
	}
}

func GetAllAttendance(c *gin.Context) {
	attendances, err, status := services.GetAllAttendance(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(status, attendances)
	}
}

func GetUserAttendance(c *gin.Context) {
	attendance, err, status := services.GetUserAttendance(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(status, attendance)
	}
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

func DeleteAttendance(c *gin.Context) {
	err, status := services.DeleteAttendance(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.Status(200)
	}
}