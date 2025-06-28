package controller

import (
	"KYC/iternals/models"
	"KYC/iternals/service"
	"KYC/utils"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	login := c.Query("login")
	password := c.Query("password")

	var u = models.User{
		Login:    login,
		Password: password,
	}

	result, err := service.CreateUser(u)
	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "data": result})
}

func SignIn(c *gin.Context) {
	userLogin := c.Query("login")
	userPassword := c.Query("password")
	result, err := service.GetUserByLogin(userLogin)
	if err != nil {
		c.JSON(404, gin.H{"status_code": 404, "error_message": "User not found"})
		return
	}

	checkPassword := service.ComparePasswords(userPassword, result.Password)
	if !checkPassword {
		c.JSON(401, gin.H{"status_code": 401, "error_message": "Wrong login or password"})
		return
	}

	token, err := utils.GenerateToken(int(result.ID), result.Login)
	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": "Error during singin proccess"})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "access_token": token, "token_type": "Bearer"})
}
