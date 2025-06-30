package controller

import (
	"KYC/iternals/models"
	"KYC/iternals/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// Update Documents godoc
// @Summary Upload user documents
// @Description Upload front | back | selfie ID photo
// @Tags documents
// @Title update
// @Accept multipart/form-data
// @Security ApiKeyAuth
// @Produce json
// @Param front formData file true "Front side of ID"
// @Param back formData file true "Back side of ID"
// @Param selfie formData file true "Selfie side if ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /documents/update [post]
func UpdateDocuments(c *gin.Context) {
	frontFile, _ := c.FormFile("front")
	backFile, _ := c.FormFile("back")
	selfieFile, _ := c.FormFile("selfie")

	frontFile.Filename = "frontID.png"
	backFile.Filename = "backID.png"
	selfieFile.Filename = "selfieID.png"

	userID := c.GetInt(userIDCtx)

	u, err := service.GetUserByID(userID)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"status_code":   401,
			"error_message": err.Error(),
		})

		return
	}

	filePath := fmt.Sprintf("./files/users/%s/docs/", u.Login)

	_ = os.MkdirAll(filePath, os.ModePerm)

	if err := c.SaveUploadedFile(frontFile, filePath+frontFile.Filename); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})

		return
	}

	if err := c.SaveUploadedFile(backFile, filePath+backFile.Filename); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})

		return
	}

	if err := c.SaveUploadedFile(selfieFile, filePath+selfieFile.Filename); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})

		return
	}

	d := models.Documents{
		Front:  frontFile.Filename,
		Back:   backFile.Filename,
		Selfie: selfieFile.Filename,
	}

	p, err := service.GetProfileByUserID(userID)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"status_code":   401,
			"error_message": err.Error(),
		})

		return
	}

	result, err := service.UpdateProfileDocuments(int(p.ID), d)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code":   400,
			"error_message": err.Error(),
		})

		return
	}

	_, err = service.UpdateUserIdentification(userID)
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
