package v1

import (
	"net/http"
	"server/model"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	if err := service.CreateCategory(data); err != nil {
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

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.DeleteCategory(id)
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
func UpdateCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	err := service.UpdateCategory(data)
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

// 分页获取
func GetCategoryList(c *gin.Context) {
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
	err, list, total := service.GetCategoryList(pageSize, pageNum)
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

// 查询单个
func GetCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	err, cate := service.GetCategory(category.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "查询失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": cate,
		})
	}
}
