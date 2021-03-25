package v1

import (
	"net/http"
	"server/middleware"
	"server/model"
	"server/service"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	err, user := service.CheckLogin(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户名或密码错误",
		})
	} else {
		token := tokenNext(c, *user)
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"token":   token,
			"message": err,
		})
	}
}

func tokenNext(c *gin.Context, user model.User) string {
	j := middleware.JWT{SigningKey: []byte(viper.GetString("jwt.signing-key"))} // 唯一签名
	claims := middleware.CustomClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                               // 签名生效时间
			ExpiresAt: time.Now().Unix() + viper.GetInt64("jwt.expires-time"), // 过期时间 7天  配置文件
			Issuer:    "ginBlog",                                              // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取token失败",
		})
	}
	return token
}
