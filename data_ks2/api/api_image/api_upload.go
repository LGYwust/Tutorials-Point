package api_image

import (
	"data_ks2/service/service_image"
	"data_ks2/utils/hash"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var ImageWhiteList = []string{".jpg", ".png", ".jpeg", ".gif", ".webp", ".svg"}

// 用户头像
func ApiIamgeUpload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "图片上传失败",
			"error": "图片上传失败",
		})
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.MapClaims)
	payload := claims["payload"].(map[string]interface{})
	userID := payload["userID"].(float64)
	savePath := path.Join("upload/user_avatar", file.Filename)
	//白名单判断
	if InImageWhiteList(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "图片上传失败",
			"error": "不支持上传该文件",
		})
		return
	}
	//重复文件判断
	filedate, _ := file.Open()
	fileHash := hash.MD5(filedate)
	filepath := service_image.CheckImage(fileHash)
	if filepath == "" {
		err := c.SaveUploadedFile(file, savePath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "图片上传失败",
				"err": "文件保存失败",
			})
			return
		}
	} else {
		savePath = filepath
	}
	if service_image.SaveImage(userID, file, savePath, fileHash) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "图片上传失败",
			"err": "文件保存失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":        "上传成功",
		"image_path": savePath,
	})
}

// 电影封面
func ApiMovieUpload(c *gin.Context) {
	// 接收电影名称
	movieName := c.PostForm("name")
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "图片上传失败",
			"error": "图片上传失败",
		})
		return
	}
	savePath := path.Join("upload/movies", file.Filename)
	//白名单判断
	if InImageWhiteList(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "图片上传失败",
			"error": "不支持上传该文件",
		})
		return
	}
	//重复文件判断
	filedate, _ := file.Open()
	fileHash := hash.MD5(filedate)
	filepath := service_image.CheckImage(fileHash)
	if filepath == "" {
		err := c.SaveUploadedFile(file, savePath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "图片上传失败",
				"err": "文件保存失败",
			})
			return
		}
	} else {
		savePath = filepath
	}
	if service_image.SaveMovieImage(movieName, savePath) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "图片上传失败",
			"err": "文件保存失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":        "上传成功",
		"image_path": savePath,
	})
}

// 电影咨询封面
func ApiConsultationUpload(c *gin.Context) {
	// 接收电影名称
	movieName := c.PostForm("name")
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "图片上传失败",
			"error": "图片上传失败",
		})
		return
	}
	savePath := path.Join("upload/consultations", file.Filename)
	//白名单判断
	if InImageWhiteList(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "图片上传失败",
			"error": "不支持上传该文件",
		})
		return
	}
	//重复文件判断
	filedate, _ := file.Open()
	fileHash := hash.MD5(filedate)
	filepath := service_image.CheckImage(fileHash)
	if filepath == "" {
		err := c.SaveUploadedFile(file, savePath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "图片上传失败",
				"err": "文件保存失败",
			})
			return
		}
	} else {
		savePath = filepath
	}
	if service_image.SaveConsultationImage(movieName, savePath) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "图片上传失败",
			"err": "文件保存失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":        "上传成功",
		"image_path": savePath,
	})
}

func InImageWhiteList(filename string) bool {
	ext := filepath.Ext(filename)
	for _, e := range ImageWhiteList {
		if e == ext {
			return false
		}
	}
	return true
}
