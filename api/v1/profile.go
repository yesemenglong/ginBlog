package v1

import (
	"net/http"
	"server/model"
	"server/service"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	var data model.Profile
	_ = c.ShouldBindJSON(&data)
	profile, err := service.GetProfile(data.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"data":    profile,
			"message": "获取失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    profile,
			"message": "获取成功",
		})
	}
}

func UpdateProfile(c *gin.Context) {
	var data model.Profile
	_ = c.ShouldBindJSON(&data)
	err := service.UpdateProfile(data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "更新失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "更新成功",
		})
	}
}
