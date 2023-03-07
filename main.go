package main

import (
	"attendance-payroll-app/initializers"
	"attendance-payroll-app/models"

	routes "attendance-payroll-app/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Announcement{})
	initializers.DB.AutoMigrate(&models.Attendance{})
	initializers.DB.AutoMigrate(&models.Leave{})
	initializers.DB.AutoMigrate(&models.Payroll{})
	initializers.DB.AutoMigrate(&models.Reimburse{})

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	routes.Routes(r)
	r.Run()
}
