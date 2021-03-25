package router

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("category")
	{
		CategoryRouter.POST("createCategory", v1.CreateCategory)
		CategoryRouter.GET("getCategoryList", v1.GetCategoryList)
		CategoryRouter.POST("updateCategory", v1.UpdateCategory)
		CategoryRouter.POST("getCategory", v1.GetCategory)
		CategoryRouter.DELETE("deleteCategory/:id", v1.DeleteCategory)
	}
}
