package middleware

import (
	"data_ks2/utils/jwts"

	"github.com/gin-gonic/gin"
)

func JwtAdmin(c *gin.Context) {
	token := c.Request.Header.Get("Token")
	if token == "" {
		c.JSON(401, gin.H{
			"msg":   "token错误",
			"error": "用户未登录",
		})
		c.Abort()
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		c.JSON(402, gin.H{
			"msg":   "token错误",
			"error": "服务器异常",
		})
		c.Abort()
		return
	}
	c.Set("claims", claims)
}
