package v1

import (
	"net/http"
	"server/service"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, err := service.UploadFile(file, fileSize)
	data := make([]string, 1)
	data[0] = url
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err,
			"url":     url,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "上传成功",
			"errno":   0,
			"data":    data,
			"url":     url,
		})
	}
}
