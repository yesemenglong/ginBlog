package router

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUploadRouter(Router *gin.RouterGroup) {
	UploadRouter := Router.Group("upload")
	{
		UploadRouter.POST("uploadFile", v1.Upload)
	}
}
