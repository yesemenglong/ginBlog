package v1

import (
	"net/http"
	"server/model"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 创建
func CreateComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)

	err := service.CreateComment(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "创建失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    data,
			"message": "创建成功",
		})
	}
}

// 查询单个
func FindComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	comment, err := service.FindComment(data.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "查询失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    comment,
			"message": "查询成功",
		})
	}
}

// 删除评论
func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.DeleteComment(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "删除失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "删除成功",
		})
	}
}

// 获取评论数量
func GetCommentCount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, total := service.GetCommentCount(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"total":   total,
			"message": "获取成功",
		})
	}
}

// 查询所有评论列表
func GetCommentList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	err, list, total := service.GetCommentList(pageSize, pageNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"list":    list,
			"total":   total,
			"message": "获取成功",
		})
	}
}

// 当前页面的评论列表
func GetCommentListFront(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	err, list, total := service.GetCommentListFront(id, pageSize, pageNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"list":    list,
			"total":   total,
			"message": "获取成功",
		})
	}
}

// 通过评论
func CheckComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))

	err := service.CheckComment(id, &data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "成功",
		})
	}
}

// 撤销评论
func UnCheckComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.UncheckComment(id, &data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "成功",
		})
	}
}
