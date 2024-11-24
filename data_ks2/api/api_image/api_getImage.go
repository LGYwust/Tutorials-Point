package api_image

import (
	"data_ks2/service/service_image"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ApiGetImage(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)
	payload := claims["payload"].(map[string]interface{})
	userID := payload["userID"].(float64)
	path := service_image.GetImage(userID)
	if path != "" {
		c.JSON(http.StatusOK, gin.H{
			"path": path,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "头像加载错误",
	})
}
