package api_user

import (
	"data_ks2/models"
	"data_ks2/service/service_user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiLogin(c *gin.Context) {
	var ul models.UserModel
	err := c.ShouldBindJSON(&ul)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "服务器异常",
		})
		return
	}
	token, flag := service_user.LoginAccount(ul.RoleID, ul.Account, ul.Password)
	if flag {
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":   "token生成失败",
				"error": "服务器错误",
			})
			return
		}
		// 将token添加到响应头中
		c.JSON(http.StatusOK, gin.H{
			"msg":   "登录成功",
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg":   "登录失败",
			"error": "用户名、密码或身份错误",
		})
	}
}
