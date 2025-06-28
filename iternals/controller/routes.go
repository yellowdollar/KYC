package controller

import (
	"github.com/gin-gonic/gin"
)

func RunServer() error {
	router := gin.Default()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/sign-up", SignUp)
		authGroup.POST("/sign-in", SignIn)
	}

	infoGroup := router.Group("/verify")
	{
		infoGroup.PUT("/info", UpdateAccountInfo)
	}

	adminGroup := router.Group("/admin")
	{
		adminGroup.POST("/create", CreateAdmin)
		adminGroup.POST("/sign-in", AdminSingIn)
	}

	router.Run(":8080")

	return nil
}
