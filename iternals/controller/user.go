package controller

import (
	"KYC/iternals/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserInformation godoc
// @Tags user
// @title GetUserInfo
// @description get user information endpoint
// @Security ApiKeyAuth
// @Access json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /user/info [get]
func GetInfo(c *gin.Context) {
	userID := c.GetInt(userIDCtx)

	u, err := service.GetUserByID(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status_code":   http.StatusNotFound,
			"error_message": err.Error(),
		})

		return
	}

	p, err := service.GetProfileByUserID(int(u.ID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status_code":   http.StatusNotFound,
			"error_message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"user":        u,
		"profile":     p,
	})
}
