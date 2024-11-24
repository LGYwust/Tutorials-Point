package api_user

import (
	"data_ks2/models"
	"data_ks2/service/service_user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Saveuserinfo(c *gin.Context) {
	var ui models.UserinfoModel
	err := c.ShouldBindJSON(&ui)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "上传失败",
			"error": "个人信息上传失败",
		})
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)
	payload := claims["payload"].(map[string]interface{})
	userID := payload["userID"].(float64)
	ui.UserID = int(userID)
	if !service_user.SaveUserinfo(ui) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "上传失败",
			"error": "个人信息上传失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "上传成功",
		"error": "个人信息上传成功",
	})
}

func Getuserinfo(c *gin.Context) {
	var ui models.UserinfoModel
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)
	payload := claims["payload"].(map[string]interface{})
	userID := payload["userID"].(float64)
	ui = service_user.Getuserinfo(int(userID))
	if ui.UserID != 0 {
		// 如果用户信息存在，则返回用户信息
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    ui,
		})
	} else {
		// 如果用户信息不存在，则返回错误
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "获取失败",
			"error": "信息获取错误",
		})
	}
}

func Getalluserinfo(c *gin.Context) {
	var um []models.UserModel
	um = service_user.Getalluser()
	if um == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "获取失败",
			"error": "信息获取错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取信息成功",
		"data": um,
	})
}
