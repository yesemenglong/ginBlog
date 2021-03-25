package router

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitCommentRouter(Router *gin.RouterGroup) {
	CommentRouter := Router.Group("comment")
	{
		CommentRouter.POST("createComment", v1.CreateComment)
		CommentRouter.POST("findComment", v1.FindComment)
		CommentRouter.DELETE("deleteComment", v1.DeleteComment)
		CommentRouter.GET("getCommentCount/:id", v1.GetCommentCount)
		CommentRouter.GET("getCommentList", v1.GetCommentList)
		CommentRouter.GET("getCommentFront/:id", v1.GetCommentListFront)
		CommentRouter.GET("checkComment/:id", v1.CheckComment)
		CommentRouter.GET("UncheckComment/:id", v1.UnCheckComment)
	}
}
