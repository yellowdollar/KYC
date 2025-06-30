package controller

import (
	"KYC/iternals/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllUsersByAdmin godoc
// @Tags admin
// @Security ApiKeyAuth
// @title Get All Users
// @Access json
// @Produce json
// @Param filter_status query string false "Filter Status"
// @Param filter_user_id query string false "Filter user id"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /admin/users [get]
func GetUsers(c *gin.Context) {
	userID := c.GetInt(userIDCtx)

	isAdmin, err := service.CheckRole(userID)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"status_code":   http.StatusBadRequest,
			"error_message": err.Error(),
		})

		return
	}

	if !isAdmin {
		c.AbortWithStatusJSON(401, gin.H{
			"status_code":   http.StatusUnauthorized,
			"error_message": "Unauthorized",
		})

		return
	}

	filterStatus := c.Query("filter_status")
	filterUserID := c.Query("filter_user_id")

	users, err := service.GetUsers(filterUserID, filterStatus)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"status_code":   http.StatusUnauthorized,
			"error_message": "Unauthorized",
		})

		return
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"data":        users,
	})
}

// ConfirmUserIdentification godoc
// @Tags admin
// @title Confirm User Identification
// @description Confirm endpoint
// @Security ApiKeyAuth
// @Access json
// @Produce json
// @Param user_id query int false "user_id"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /admin/confirm [put]
func ConfirmIdentification(c *gin.Context) {
	userID := c.Query("user_id")
	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})
	}

	adminID := c.GetInt(userIDCtx)
	isAdmin, err := service.CheckRole(adminID)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"status_code":   http.StatusUnauthorized,
			"error_message": err.Error(),
		})

		return
	}

	if !isAdmin {
		c.AbortWithStatusJSON(401, gin.H{
			"status_code":   http.StatusUnauthorized,
			"error_message": "Unauthorized",
		})

		return
	}

	result, err := service.ConfirmUserIdentifyStatus(userIDint)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code":   http.StatusBadRequest,
			"error_message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"message":     "User identified successfully",
		"data":        result,
	})

}
