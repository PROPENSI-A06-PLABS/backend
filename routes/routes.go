package routes

import (
	"attendance-payroll-app/controllers"
	"attendance-payroll-app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)

	userRouter := router.Group("/user")
	{
		userRouter.GET("/retrieve-users", controllers.RetrieveUsers)
		userRouter.GET("/retrieve-user/:id", controllers.RetrieveUser)
		userRouter.POST("/create-user", controllers.CreateUser)
		userRouter.DELETE("/delete-user/:id", controllers.DeleteUser)
		userRouter.PUT("/update-user/:id", controllers.UpdateUser)
		userRouter.PUT("/change-status/:id", controllers.ChangeStatus)
		userRouter.POST("/delete-users", controllers.DeleteUsers)
		userRouter.POST("/activate-users", controllers.ActivateUsers)
		userRouter.POST("/deactivate-users", controllers.DeactivateUsers)
	}

	dashboardRouter := router.Group("/dashboard")
	{
		dashboardRouter.POST("/check-in", controllers.CheckIn)
		dashboardRouter.POST("/check-out/:id", controllers.CheckOut)
	}

	attendanceRouter := router.Group("/attendance")
	{
		attendanceRouter.PUT("/update-attendance/:id",  controllers.UpdateAttendance)
	}

	router.POST("/login", controllers.Login)
	router.POST("/logout", middleware.RequireAuth, controllers.Logout)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.GET("/get-attendance", controllers.GetAllAttendance)
	router.GET("/get-attendance/:id", controllers.GetUserAttendance)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	},
	)
}
