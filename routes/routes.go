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
	router.GET("/retrive-users", controllers.RetriveUsers)
	router.POST("/create-user", controllers.CreateUser)
	router.DELETE("/delete-user/:id", controllers.DeleteUser)
	router.PUT("/update-user/:id", controllers.UpdateUser)
	// router.GET("/retrive-user/:id", controllers.)


	router.POST("/login", controllers.Login)
	router.POST("/logout", middleware.RequireAuth, controllers.Logout)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	},
	)
}
