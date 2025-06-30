package controller

import (
	"KYC/iternals/models"
	"KYC/iternals/service"
	"KYC/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @Tags Auth
// @title SignUp endpoint
// @version v1.0
// @description Login & password registration
// @Accept json
// @Produce json
// @Param userSignUp body models.UserSignUp true "User data"
// @Success 201 {string} User Created Successfully
// @Failure 400 {object} map[string]string
// @Router /auth/sign-up [post]
func SignUp(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})

		return
	}

	if result, err := service.CreateUser(u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})

		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status_code": 201,
			"data":        result,
		})
	}
}

// SignIn godoc
// @Tags Auth
// @title SignIn
// @version v1.0
// @description Sign In endpoint
// @Access json
// @Produce json
// @Param userSignIn body models.UserSignIn true "User Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/sign-in [post]
func SignIn(c *gin.Context) {
	var u models.UserSignIn

	if err := c.ShouldBindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})

		return
	}

	user, err := service.GetUserByLogin(u.Login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status_code":   404,
			"error_message": "Wrong login or Password",
		})

		return
	}

	token, err := utils.GenerateToken(int(user.ID), user.Login, user.Role)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"token_type":   "Bearer",
	})
}
