package router

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitProfileRouter(Router *gin.RouterGroup) {
	ProfileRouter := Router.Group("profile")
	{
		ProfileRouter.POST("getProfile", v1.GetProfile)
		ProfileRouter.POST("updateProfile", v1.UpdateProfile)
	}
}
