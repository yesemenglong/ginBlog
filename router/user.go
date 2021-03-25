package router

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("createUser", v1.CreateUser)
		UserRouter.GET("getUserList", v1.GetUserList)
		UserRouter.POST("updateUser", v1.EditUser)
		UserRouter.POST("getUser", v1.GetUser)
		UserRouter.DELETE("deleteUser/:id", v1.DeleteUser)
		UserRouter.POST("login", v1.Login)
		UserRouter.POST("ChangePassword", v1.ChangePassword)
	}
}
