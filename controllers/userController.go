package controllers

import (
	"attendance-payroll-app/initializers"
	"attendance-payroll-app/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	// get data from body
	var body struct {
		Name         string
		Email        string
		Phone        int
		Password     string
		Position     string
		Status       bool
		StartWork    time.Time
		Tenure       string
		ContractType string
		GrossSalary  int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})

		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password",
		})

		return
	}

	// create user
	newUser := models.User{
		Name:         body.Name,
		Email:        body.Email,
		Phone:        body.Phone,
		Password:     string(hash),
		Position:     body.Position,
		Status:       body.Status,
		StartWork:    body.StartWork,
		Tenure:       body.Tenure,
		ContractType: body.ContractType,
		GrossSalary:  body.GrossSalary,
	}

	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})

		return
	}

	//return
	c.JSON(http.StatusOK, gin.H{
		"user": newUser,
	})
}

func Login(c *gin.Context) {
	// get email and password
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
	}

	// search user by email
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	// err := initializers.DB.First(&user, body.Email).Error //fetch error kalo not found

	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})
	}

	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"message": "not found",
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"user": user,
	// 	})
	// }

	// compare pass with saved pass
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})

		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 hari
	})

	// sign the token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create token",
		})

		return
	}
	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true) // 1hour*24*30 = 30 hari
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
}
