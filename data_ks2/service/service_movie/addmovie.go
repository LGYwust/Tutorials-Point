package service_movie

import (
	"data_ks2/global"
	"data_ks2/models"
)

func Addmovie(movie models.MovieModel) int {
	var mm models.MovieModel
	err := global.DB.Take(&mm, "name = ?", movie.Name).Error
	if err == nil {
		return 0 // 电影已存在，返回 0 表示失败
	}
	err = global.DB.Create(&movie).Error
	if err != nil {
		// 如果添加失败，返回 -1 表示数据库插入错误
		return -1
	}
	// 返回 1 表示添加成功
	return 1
}

func Deltmovie(name string) int {
	var mm models.MovieModel
	err := global.DB.Take(&mm, "name = ?", name).Error
	if err != nil {
		return 0
	}
	err = global.DB.Delete(&mm).Error
	if err != nil {
		return 0
	}
	return 1
}
