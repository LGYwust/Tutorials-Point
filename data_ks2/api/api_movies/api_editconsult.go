package api_movies

import (
	"data_ks2/models"
	"data_ks2/service/service_movie"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ApiAddconsultation(c *gin.Context) {
	var mc models.ConsultationModel
	err := c.ShouldBindJSON(&mc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "前端获取数据失败",
			"error": "前端获取数据失败",
		})
		return
	}
	flag := service_movie.Addconsultation(mc)
	if flag == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "添加失败",
			"error": "数据库插入错误",
		})
		return
	}
	if flag == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "添加失败",
			"error": "电影已存在，请联系管理员",
		})
		return
	}
	if flag == 1 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "添加成功",
		})
		return
	}
}

func ApiDelconsultation(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "获取信息失败",
			"error": "无法读取请求体",
		})
		return
	}
	title := string(body) // 将请求体的字节数据转换为字符串
	title = strings.ReplaceAll(title, `"`, "")
	// 后续的逻辑
	flag := service_movie.Deltconsultations(title)
	if flag == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "删除失败",
			"error": "删除信息失败",
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
