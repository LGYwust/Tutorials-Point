package service_image

import (
	"data_ks2/global"
	"data_ks2/models"
	"fmt"
)

// 观看人数最多的20部电影
func Gethotmovies() []models.MovieModel {
	var moviesList []models.MovieModel
	// 从数据库中查询观看人数最多的20部电影
	result := global.DB.Order("people DESC").Limit(20).Find(&moviesList)
	if result.Error != nil {
		// 处理错误，例如记录日志或返回默认值
		return nil
	}
	for i := range moviesList {
		moviesList[i].Image = fmt.Sprint("http://", global.Config.System.IP, ":", global.Config.System.Port, "/", moviesList[i].Image)
	}
	return moviesList
}

// 推荐电影
func Getrecommendmovies() []models.MovieModel {
	var moviesList []models.MovieModel
	result := global.DB.Order("RAND()").Limit(3).Find(&moviesList)
	if result.Error != nil {
		// 处理错误，例如记录日志或返回默认值
		return nil
	}
	for i := range moviesList {
		moviesList[i].Image = fmt.Sprint("http://", global.Config.System.IP, ":", global.Config.System.Port, "/", moviesList[i].Image)
	}
	return moviesList
}

// 搜索电影
func Getsearchmovies(data string) models.MovieModel {
	var movie models.MovieModel
	global.DB.Take(&movie, "name = ?", data)
	return movie
}

// 所有电影
func Getallmovies() []models.MovieModel {
	var moviesList []models.MovieModel
	result := global.DB.Find(&moviesList)
	if result.Error != nil {
		return nil
	}
	for i := range moviesList {
		moviesList[i].Image = fmt.Sprint("http://", global.Config.System.IP, ":", global.Config.System.Port, "/", moviesList[i].Image)
	}
	return moviesList
}

// 筛选电影
func Getselectmovies(category string, area string, hon string) []models.MovieModel {
	var moviesList []models.MovieModel
	// 使用 LIKE 操作符构建查询，匹配空格分隔的类别和地区字段中的任意值
	query := global.DB
	// 根据是否传入 category 进行条件过滤
	if category != "全部" {
		if category == "其他" {
			// 排除指定国家的电影
			excludedCategories := []string{"动作", "悬疑", "科幻", "剧情", "儿童", "惊悚", "爱情", "喜剧"}
			for _, category := range excludedCategories {
				query = query.Where("category NOT LIKE ?", "%"+category+"%")
			}
		} else {
			query = query.Where("category LIKE ?", "%"+category+"%")
		}
	}
	// 根据是否传入 area 进行条件过滤
	if area != "全部" {
		if area == "其他" {
			// 排除指定国家的电影
			excludedCountries := []string{"英国", "美国", "中国", "日本", "韩国", "德国", "瑞士", "印度"}
			for _, country := range excludedCountries {
				query = query.Where("country NOT LIKE ?", "%"+country+"%")
			}
		} else {
			query = query.Where("country LIKE ?", "%"+area+"%")
		}
	}
	// 根据 hon 的值来决定排序方式
	if hon == "Hot" {
		query = query.Order("people DESC")
	} else if hon == "New" {
		query = query.Order("time DESC")
	} else {
		// 如果 hon 不是 "Hot" 或 "New"，可以设置一个默认的排序方式
		query = query.Order("people DESC")
	}
	// 执行查询并将结果存储在 moviesList 中
	result := query.Find(&moviesList)
	if result.Error != nil {
		return nil
	}
	for i := range moviesList {
		moviesList[i].Image = fmt.Sprint("http://", global.Config.System.IP, ":", global.Config.System.Port, "/", moviesList[i].Image)
	}
	return moviesList
}
