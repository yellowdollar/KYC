package controller

import (
	"KYC/iternals/models"
	"KYC/iternals/service"

	"github.com/gin-gonic/gin"
)

// UpdateUserInfo godoc
// @Tags profile
// @Security ApiKeyAuth
// @title Update
// @version v1.0
// @description UpdateUserInfo endpoint
// @Access json
// @Produce json
// @Param profileUpdate body models.ProfileUpdate true "Profile data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /profile/update [put]
func UpdateUserInfo(c *gin.Context) {
	var p models.Profile

	if err := c.ShouldBindJSON(&p); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})

		return
	}

	userID := c.GetInt(userIDCtx)

	result, err := service.UpdateProfileInfo(userID, p)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"data":        result,
	})
}

// GetUserInfo godoc
// @Tags profile
// @Security ApiKeyAuth
// @title Get
// @version v1.0
// @description GetUserInfo endpoint
// @Access json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /profile/get [get]
func GetUserInfoByID(c *gin.Context) {
	userID := c.GetInt(userIDCtx)

	result, err := service.GetProfileByUserID(userID)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"status_code":   404,
			"error_message": "user's profile not found",
		})

		return
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"data":        result,
	})
}
