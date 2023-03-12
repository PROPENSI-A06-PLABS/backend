package services

import (
	"attendance-payroll-app/initializers"
	"attendance-payroll-app/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/dongri/phonenumber"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func DeactivateUsers(c *gin.Context) ([]models.User, error, int) {
	type UserId struct{
		ID int
	}
	users := []models.User{}

	// get all input
	respBody , err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return users, err, http.StatusBadRequest
	}

	// convert json list to dictionary
	var body []map[string]int
	err = json.Unmarshal(respBody, &body)
	if err != nil {
		return users, err, http.StatusBadRequest
	}


	for _ , value := range body{
		user := models.User{}
		// get user by id
		err = initializers.DB.First(&user, value["Id"]).Error
		if err != nil {
			return users, err, http.StatusBadRequest
		}

		// update user status (false)
		status := false
		err = initializers.DB.Model(&user).Update("status", status).Error
		if err != nil {
			return users, err, http.StatusBadRequest
		}
		user.Status = status
		users = append(users, user)
	}
	return users, nil, http.StatusOK
}

func ActivateUsers(c *gin.Context) ([]models.User, error, int) {
	type UserId struct{
		ID int
	}
	users := []models.User{}

	// get all input
	respBody , err := ioutil.ReadAll(c.Request.Body)
	if (err != nil){
		return users, err, http.StatusBadRequest
	}

	// convert json list to dictionary
	var body []map[string]int
	err = json.Unmarshal(respBody, &body)
	if (err != nil){
		return users, err, http.StatusBadRequest
	}

	for _ , value := range body{
		user := models.User{}
		// get user by id
		err = initializers.DB.First(&user, value["Id"]).Error
		if (err != nil){
			return users, err, http.StatusBadRequest
		}

		// update user status (true)
		status := true
		err = initializers.DB.Model(&user).Update("status", status).Error
		if (err != nil){
			return users, err, http.StatusBadRequest
		}
		user.Status = status
		users = append(users, user)
	}
	return users, nil, http.StatusOK
}

func DeleteUsers(c *gin.Context) (error, int){
	type UserId struct{
		ID int
	}

	// get all input
	respBody , err := ioutil.ReadAll(c.Request.Body)
	if (err != nil){
		return err, http.StatusBadRequest
	}

	// convert json list to dictionary
	var body []map[string]int
	err = json.Unmarshal(respBody, &body)
	if (err != nil){
		return err, http.StatusBadRequest
	}

	for _, value := range body{
		userModel := models.User{}
		err = initializers.DB.Delete(&userModel, value["Id"]).Error
		if (err != nil){
			return err, http.StatusBadRequest
		}
	}
	return nil, http.StatusOK
}

func ChangeStatus(c *gin.Context) (models.User,error,int) {
	// get user by id
	id := c.Param("id")
	user := models.User{}
	err := initializers.DB.First(&user, id).Error
	if (err != nil){
		return user, err, http.StatusBadRequest
	}
	// update user status (negasi)
	status := user.Status
	err = initializers.DB.Model(&user).Update("status", !status).Error
	if (err != nil){
		return user, err, http.StatusBadRequest
	}

	user.Status = !status
	return user, err, http.StatusOK
}


func UpdateUser(c *gin.Context) (models.User, error, int) {
	// get user by id
	id := c.Param("id")
	user := models.User{}
	err := initializers.DB.First(&user, id).Error
	if (err != nil){
		return user, err, http.StatusBadRequest
	}
	// get user's input
	input := models.User{}
	err = c.Bind(&input)
	if (err != nil){
		return user, err, http.StatusBadRequest
	}

	// validate input
	validate := validator.New()
	err = validate.Struct(input)
	if (err != nil){
		return user, err, http.StatusBadRequest
	}

	// update user
	err = initializers.DB.Model(&user).Updates(
		models.User{
			Name: input.Name,
			Phone: phonenumber.Parse(input.Phone, "IDN"),
			Email: input.Email,
			BornDay: input.BornDay,
			AddressDetail: input.AddressDetail,
			IdentityNumber: input.IdentityNumber,
			AccountNumber: input.AccountNumber,
			NPWP: input.NPWP,
			KPJ: input.KPJ,
			Github: input.Github,
			BankName: input.BankName,
			JKN_KIS: input.JKN_KIS,
			Gitlab: input.Gitlab,
			ExtraInfo: input.ExtraInfo,
			PLABSMail: input.PLABSMail,
			Position: input.Position,
			GrossSalary: input.GrossSalary,
			Role: input.Role,
			ContractType: input.ContractType,
			Password: input.Password,
			StartWork: input.StartWork,
			Status: input.Status,
			NPWPDocument: input.NPWPDocument,
			KTPDocument: input.KTPDocument,
			CVDocument: input.CVDocument,
			ContractDocument: input.ContractDocument,
			ProfilePhoto: input.ProfilePhoto,
		},
	).Error
	if (err != nil){
		return user, err, http.StatusBadRequest
	}
	input.Id = user.Id
	return input, nil, http.StatusOK
}

func DeleteUser(c *gin.Context) (error, int) {
	// delete user by id
	id := c.Param("id")
	user := models.User{}
	err := initializers.DB.Delete(&user, id).Error
	if err != nil{
		return err, http.StatusBadRequest
	}
	return nil, http.StatusOK
}

func RetrieveUsers(c *gin.Context) ([]models.User, error, int) {
	// get all users
	users := []models.User{}
	err := initializers.DB.Find(&users).Error
	if (err != nil){
		return users, err, http.StatusBadRequest
	}

	return users, nil, http.StatusOK
}

func RetrieveUser(c *gin.Context) (models.User, error, int) {
	// find user by id
	id := c.Param("id")
	user := models.User{}
	err := initializers.DB.Find(&user, id).Error
	if (err != nil){
		return user, err, http.StatusBadRequest
	}

	return user, nil, http.StatusOK
}

func CreateUser(c *gin.Context) (models.User, string, int) {
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