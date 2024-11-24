package api_user

import (
	"data_ks2/models"
	"data_ks2/service/service_user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRegister(c *gin.Context) {
	var ur models.UserModel
	err := c.ShouldBindJSON(&ur)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if service_user.CheckUserName(ur.Account) {
		c.JSON(502, gin.H{
			"msg":   "用户创建失败",
			"error": "用户名已存在",
		})
		return
	}
	if service_user.AddUser(ur.Account, ur.Email, ur.Phone, ur.Password) {
		c.JSON(502, gin.H{
			"msg":   "用户创建失败",
			"error": "数据库创建失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":   "用户创建成功",
		"error": "",
		"data":  1,
	})
}
