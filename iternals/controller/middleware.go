package controller

import (
	"KYC/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userIDCtx  = "userID"
)

func CheckUserAuth(c *gin.Context) {
	header := c.GetHeader(authHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error_message": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error_message": "invalid auth header",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error_message": "token row is empty",
		})
		return
	}

	accessToken := headerParts[1]

	claims, err := utils.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	c.Set(userIDCtx, claims.UserID)
	c.Next()
}
