package routes

import (
	"attendance-payroll-app/controllers"
	"attendance-payroll-app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	// CRUD user
	router.GET("/retrieve-users", middleware.RequireAuth, controllers.RetrieveUsers)
	router.POST("/create-user", middleware.RequireAuth, controllers.CreateUser)
	router.DELETE("/delete-user/:id", middleware.RequireAuth, controllers.DeleteUser)
	router.PUT("/update-user/:id", middleware.RequireAuth, controllers.UpdateUser)

	router.POST("/login", controllers.Login)
	router.POST("/logout", middleware.RequireAuth, controllers.Logout)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.POST("/create-attendance", controllers.CheckIn)
	router.GET("/get-attendance", controllers.GetAllAttendance)
	router.GET("/get-attendance/:id", controllers.GetUserAttendance)
	// router.POST("/create-attendance", controllers.CheckIn)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	},
	)
}
