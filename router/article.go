package router

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitArticleRouter(Router *gin.RouterGroup) {
	ArticleRouter := Router.Group("article")
	{
		ArticleRouter.POST("createArticle", v1.CreateArticle)
		ArticleRouter.GET("getArticleList", v1.GetArticleList)
		ArticleRouter.POST("updateArticle", v1.UpdateArticle)
		ArticleRouter.DELETE("deleteArticle/:id", v1.DeleteArticle)
		ArticleRouter.POST("findArticle", v1.FindArticle)
		ArticleRouter.GET("getCateArticleList/:id", v1.GetCateArticleList)
	}
}
