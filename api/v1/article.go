package v1

import (
	"net/http"
	"server/model"
	"server/model/request"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	if err := service.CreateArticle(data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"data":    data,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    data,
			"message": "创建成功",
		})
	}
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "删除成功",
		})
	}
}

// 修改分类
func UpdateArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	err := service.UpdateArticle(data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "更新成功",
		})
	}
}

func GetArticleList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	title := c.Query("title")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	err, list, total := service.GetArticleList(pageSize, pageNum, title)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"data":    list,
			"total":   total,
			"message": "获取失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    list,
			"total":   total,
			"message": "获取成功",
		})
	}
}

func FindArticle(c *gin.Context) {
	var a request.ArticleRequest
	_ = c.ShouldBindJSON(&a)
	id, _ := strconv.Atoi(a.Id)
	err, data := service.FindArticle(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"data":    data,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    data,
			"message": "获取成功",
		})
	}
}

func GetCateArticleList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	id, _ := strconv.Atoi(c.Param("id"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	err, list, total := service.GetCateArticleList(id, pageSize, pageNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"data":    list,
			"total":   total,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    list,
			"total":   total,
			"message": "获取成功",
		})
	}
}
