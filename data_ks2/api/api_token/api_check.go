package api_token

import (
	"data_ks2/utils/jwts"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tokendate struct {
	token string
}

func CheckToken(c *gin.Context) {
	var tokendate tokendate
	err := c.ShouldBindJSON(&tokendate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "服务器异常",
			"flag":  false,
		})
		return
	}
	_, err = jwts.ParseToken(tokendate.token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户认证失效",
			"flag":  false,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  "用户认证成功",
		"flag": true,
	})
}
