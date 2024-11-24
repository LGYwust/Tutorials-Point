package api_user

import (
	"data_ks2/service/service_user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ApiEdituser(c *gin.Context) {
	type userdata struct {
		Account string
		Email   string
		Phone   string
		RoleID  string
	}
	var um userdata
	err := c.ShouldBindJSON(&um)
	urid, err := strconv.Atoi(um.RoleID)
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)
	payload := claims["payload"].(map[string]interface{})
	RoleID := payload["roleID"].(float64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "获取信息失败",
			"error": "获取信息失败",
		})
		return
	}
	flag := service_user.EditUser(um.Account, urid, int(RoleID))
	if flag == -2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "更新失败",
			"error": "不能修改自己的权限",
		})
		return
	}
	if flag == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "更新失败",
			"error": "没有该操作的权限",
		})
		return
	}
	if flag == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "更新失败",
			"error": "更新失败",
		})
		return
	}
	if flag == 1 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "更新成功",
		})
		return
	}

}

func ApiDeluser(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "获取信息失败",
			"error": "无法读取请求体",
		})
		return
	}

	account := string(body) // 将请求体的字节数据转换为字符串
	fmt.Print(account, "-------------")

	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)
	payload := claims["payload"].(map[string]interface{})
	RoleID := payload["roleID"].(float64)

	// 后续的逻辑
	flag := service_user.DeltUser(account, int(RoleID))
	if flag == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "删除失败",
			"error": "删除信息失败",
		})
		return
	}
	if flag == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "删除失败",
			"error": "无法删除管理员",
		})
		return
	}
	if flag == 1 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除成功",
		})
		return
	}
}
