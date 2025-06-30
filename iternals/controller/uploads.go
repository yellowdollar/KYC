package controller

import (
	"github.com/gin-gonic/gin"
)

// Upload godoc
// @Summary upload files
// @Tags upload
// @Produce application/octet-stream
// @Param file_path query string true "Full path to the file"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /upload [get]
func Upload(c *gin.Context) {
	filepath := c.Query("file_path")

	c.File(filepath)
}
