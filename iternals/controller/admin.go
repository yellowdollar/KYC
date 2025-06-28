package controller

import (
	"KYC/iternals/models"
	"KYC/iternals/service"
	"KYC/utils"

	"github.com/gin-gonic/gin"
)

func CreateAdmin(c *gin.Context) {
	login := c.Query("login")
	password := c.Query("password")

	var admin = models.Admin{
		Login:    login,
		Password: password,
	}

	_, err := service.CreateAdmin(admin)
	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": err})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "data": admin})
}

func AdminSingIn(c *gin.Context) {
	login := c.Query("login")
	password := c.Query("password")

	result, err := service.GetAdminByLogin(login)
	if err != nil {
		c.JSON(401, gin.H{"status_code": 404, "error_message": "wrong login or password"})
		return
	}

	if comparePassword := service.ComparePasswords(password, result.Password); !comparePassword {
		c.JSON(401, gin.H{"status_code": 404, "error_message": "wrong login or password"})
		return
	}

	token, err := utils.GenerateToken(int(result.ID), result.Login)
	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": "Error during sign in proccess"})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "access_token": token, "token_type": "Bearer"})
}
