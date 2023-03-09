package controllers

import (
	"attendance-payroll-app/initializers"
	"attendance-payroll-app/models"
	// "attendance-payroll-app/services"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func UpdateUser(c *gin.Context){
	id := c.Param("id")
	user := models.User{}
	db_user := models.User{}

	c.Bind(&user) // from fe

	initializers.DB.First(&db_user,id) // from database
	initializers.DB.Model(&db_user).Updates(user)
	
}


func DeleteUser(c *gin.Context){
	id := c.Param("id")
	user := models.User{}
	initializers.DB.Delete(&user, id)
	c.Status(200)
}

func RetriveUsers(c *gin.Context){
	// get all data users
	users := []models.User{}
	initializers.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	// get data from body
	body := models.User{}

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

	// validate input
	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// create user
	body.Password = string(hash)
	result := initializers.DB.Create(&body)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})

		return
	}

	//return
	c.JSON(http.StatusOK, gin.H{
		"user": body,
	})
}

// func CreateUser(c *gin.Context) {
// 	user, err, status := services.CreateUser(c)
// 	if (err != ""){
// 		c.JSON(status, gin.H{
// 			"message": err,
// 		})
// 	} else{
// 		c.JSON(status, user)
// 	}
// }

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
