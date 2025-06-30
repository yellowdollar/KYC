package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer() error {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/sign-up", SignUp)
		authGroup.POST("/sign-in", SignIn)
	}

	profile := router.Group("/profile", CheckUserAuth)

	profileGroup := profile.Group("/")
	{
		profileGroup.PUT("/update", UpdateUserInfo)
		profileGroup.GET("/get", GetUserInfoByID)
	}

	user := router.Group("/user", CheckUserAuth)

	userGroup := user.Group("/")
	{
		userGroup.GET("/info", GetInfo)
	}

	documents := router.Group("/documents", CheckUserAuth)

	documentsGroup := documents.Group("/")
	{
		documentsGroup.POST("/update", UpdateDocuments)
	}

	admin := router.Group("/admin", CheckUserAuth)

	adminGroup := admin.Group("/")
	{
		adminGroup.GET("/users", GetUsers)
		adminGroup.PUT("/confirm", ConfirmIdentification)
	}

	router.Run(":8080")

	return nil
}
