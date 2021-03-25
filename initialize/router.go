package initialize

import (
	"server/middleware"
	"server/router"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(viper.GetString("server.system.mode"))
	var Router = gin.Default()
	Router.Use(middleware.Cors())
	//Router.LoadHTMLGlob("../static/admin/index.html")
	//Router.Static("admin/static", "static/admin/static")
	//
	//Router.GET("/admin", func(c *gin.Context) {
	//	c.HTML(200, "index.html", nil)
	//})

	PrivateGroup := Router.Group("api/v1")
	PrivateGroup.Use(middleware.Logger())
	{
		router.InitUserRouter(PrivateGroup)
		router.InitCategoryRouter(PrivateGroup)
		router.InitArticleRouter(PrivateGroup)
		router.InitUploadRouter(PrivateGroup)
		router.InitProfileRouter(PrivateGroup)
		router.InitCommentRouter(PrivateGroup)
	}

	return Router
}
