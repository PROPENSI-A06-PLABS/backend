package services

import (
	"attendance-payroll-app/models"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"attendance-payroll-app/initializers"
	"net/http"
)

func CreateUser(c *gin.Context) (models.User, string, int){
	// get data from body
	newUser := models.User{}
	err := c.Bind(&newUser)
	if (err != nil){
		return newUser, "Failed to read body", http.StatusBadRequest
	}

	// validate input
	validate := validator.New()
	err = validate.Struct(newUser)
	if (err != nil){
		return newUser, err.Error(), http.StatusBadRequest
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if (err != nil){
		return newUser, "Failed to hash password", http.StatusBadRequest
	}
	newUser.Password = string(hash)

	// gorm (repo)
	result := initializers.DB.Create(&newUser)
	if result.Error != nil{
		return newUser, "Failed to create user", http.StatusBadRequest
	}

	//success crete new user
	return newUser, "" , http.StatusOK
}