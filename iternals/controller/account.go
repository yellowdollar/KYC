package controller

import (
	"KYC/iternals/models"
	"KYC/iternals/service"
	"KYC/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateAccountInfo(c *gin.Context) {
	token := c.Query("token")

	tokenParse, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(401, gin.H{"status_code": 401, "error_message": "unauthorized"})
		return
	}

	userID := tokenParse.UserID

	name := c.Query("name")
	surname := c.Query("surname")
	male := c.Query("male")

	ageStr := c.Query("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": "input valid age"})
		return
	}

	var a = models.Account{
		Name:    name,
		Surname: surname,
		Male:    male,
		Age:     age,
	}

	result, err := service.UpdateAccountInfo(userID, a)
	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": err})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "data": result})
}
