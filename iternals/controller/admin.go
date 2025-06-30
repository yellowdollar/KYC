package controller

import (
	"KYC/iternals/models"
	"KYC/iternals/service"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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

	uWp := models.UserWithoutProfile{
		ID:           result.ID,
		Login:        result.Login,
		Role:         result.Role,
		IsIdentified: result.IsIdentified,
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"message":     "User identified successfully",
		"data":        uWp,
	})

}

// GetUserDocsByUserID godoc
// @Summary      Get user documents by user ID
// @Description  Returns paths to user documents if they exist
// @Tags         admin
// @Security     ApiKeyAuth
// @Produce      json
// @Param        id    query     int  true  "User ID"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /admin/user-docs [get]
func GetUserDocsByUserID(c *gin.Context) {
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

	userID := c.Query("id")
	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})

		return
	}

	u, err := service.GetUserByID(userIDint)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"status_code":   404,
			"error_message": err.Error(),
		})

		return
	}

	filePath := fmt.Sprintf("./files/users/%s/docs/", u.Login)

	files := []string{"frontID.png", "backID.png", "selfieID.png"}
	result := make(map[string]string)

	for _, file := range files {
		full := filepath.Join(filePath, file)
		if _, err := os.Stat(full); err == nil {
			result[file] = full
		} else {
			result[file] = ""
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":   u.ID,
		"login":     u.Login,
		"documents": result,
	})
}
