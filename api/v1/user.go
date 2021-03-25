package v1

import (
	"net/http"
	"server/model"
	"server/model/request"
	"server/service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

// 添加用户
func CreateUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	err := utils.Validator(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "校验不通过",
			"err":     err,
		})
		return
	}
	if err := service.CreateUser(data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"data":    data,
			"message": "创建失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    data,
			"message": err,
		})
	}
}

// 查询用户列表
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	username := c.Query("username")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	err, list, total := service.GetUserList(username, pageSize, pageNum)
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
			"message": err,
		})
	}
}

// 修改用户
func EditUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	err := service.UpdateUser(data)
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

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.DeleteUser(id)
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

// 查询单个
func GetUser(c *gin.Context) {
	var u model.User
	_ = c.ShouldBindJSON(&u)
	err, user := service.GetUser(u.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": user,
		})
	}
}

// 修改密码
func ChangePassword(c *gin.Context) {
	var user request.ChangePassword
	_ = c.ShouldBindJSON(&user)
	u := model.User{Username: user.Username, Password: user.OldPassword}
	err, _ := service.ChangePassword(u, user.NewPassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "修改失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "修改成功",
		})
	}
}
