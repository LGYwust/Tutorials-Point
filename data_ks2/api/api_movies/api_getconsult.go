package api_movies

import (
	"data_ks2/models"
	"data_ks2/service/service_movie"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiGetconsult(c *gin.Context) {
	var uc []models.ConsultationModel
	uc = service_movie.Getconsultation()
	if uc == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "获取数据失败",
			"error": "获取数据失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":           "获取数据成",
		"consultations": uc,
	})
}
