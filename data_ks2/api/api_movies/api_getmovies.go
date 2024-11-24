package api_movies

import (
	"data_ks2/models"
	"data_ks2/service/service_image"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiGethotmovies(c *gin.Context) {
	moviesList := []models.MovieModel{}
	moviesList = service_image.Gethotmovies()
	if moviesList == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "后台数据拉起失败",
			"error": "电影获取失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "数据获取成功",
		"movies": moviesList,
	})
}

func ApiGetRecommendmovies(c *gin.Context) {
	moviesList := []models.MovieModel{}
	moviesList = service_image.Getrecommendmovies()
	if moviesList == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "后台数据拉起失败",
			"error": "电影获取失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "数据拉起成功",
		"movies": moviesList,
	})
}

func ApiGetsearchmovies(c *gin.Context) {
	movie := models.MovieModel{}
	var data string
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg":   "前端获取数据失败",
			"error": "获取数据失败",
		})
		return
	}
	movie = service_image.Getsearchmovies(data)
	if movie.Name != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "获取数据失败",
			"error": "没有找到此电影",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "获取数据成功",
		"movie": movie,
	})
}

func ApiGetallmovies(c *gin.Context) {
	moviesList := []models.MovieModel{}
	moviesList = service_image.Getallmovies()
	if moviesList == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "获取信息错误",
			"error": "获取数据失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "获取信息成功",
		"movies": moviesList,
	})
}

func ApiGetselectmovies(c *gin.Context) {
	moviesList := []models.MovieModel{}
	var data []string
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "前端获取数据失败",
			"error": "前端获取数据失败",
		})
		return
	}
	moviesList = service_image.Getselectmovies(data[0], data[1], data[2])
	if moviesList == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "获取信息错误",
			"error": "获取数据失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "获取信息成功",
		"movies": moviesList,
	})
}
