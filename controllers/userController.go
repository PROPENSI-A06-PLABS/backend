package controllers

import (
	"attendance-payroll-app/initializers"
	"attendance-payroll-app/models"
	"attendance-payroll-app/services"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// deactivate multiple users
func DeactivateUsers(c *gin.Context) {
	users, err, status := services.DeactivateUsers(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(status, users)
	}
}

// activate multiple users
func ActivateUsers(c *gin.Context) {
	users, err, status := services.ActivateUsers(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(status, users)
	}
}

// delete multiple users
func DeleteUsers(c *gin.Context) {
	err, status := services.DeleteUsers(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.Status(status)
	}
}

// change user status by id (negasi)
func ChangeStatus(c *gin.Context) {
	user, err, status := services.ChangeStatus(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(status, user)
	}
}

// update user attribute by id
func UpdateUser(c *gin.Context) {
	user, err, status := services.UpdateUser(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(status, user)
	}
}

// delete user by id
func DeleteUser(c *gin.Context) {
	err, status := services.DeleteUser(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.Status(200)
	}
}

// get all users
func RetrieveUsers(c *gin.Context) {
	users, err, status := services.RetrieveUsers(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(status, users)
	}

}

// get user by id
func RetrieveUser(c *gin.Context) {
	user, err, status := services.RetrieveUser(c)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(status, user)
	}
}

// create user
func CreateUser(c *gin.Context) {
	user, err, status := services.CreateUser(c)
	if err != "" {
		c.JSON(status, gin.H{
			"message": err,
		})
	} else {
		c.JSON(status, user)
	}
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
		// "exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 hari
		// "exp": time.Now().Add(time.Minute * 5).Unix(), // 5 menit
		"exp": time.Now().Add(time.Hour * 12).Unix(), // 12 jam
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
	// c.Set("Authorization", tokenString)
	// c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true) // 1hour*24*30 = 30 hari
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user":  user,
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
