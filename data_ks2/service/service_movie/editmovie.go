package service_movie

import (
	"data_ks2/global"
	"data_ks2/models"
)

func Addconsultation(consultations models.ConsultationModel) int {
	var mc models.ConsultationModel
	err := global.DB.Take(&mc, "title = ?", consultations.Title).Error
	if err == nil {
		return 0 // 电影已存在，返回 0 表示失败
	}
	err = global.DB.Create(&consultations).Error
	if err != nil {
		// 如果添加失败，返回 -1 表示数据库插入错误
		return -1
	}
	// 返回 1 表示添加成功
	return 1
}

func Deltconsultations(title string) int {
	var mc models.ConsultationModel
	err := global.DB.Take(&mc, "title = ?", title).Error
	if err != nil {
		return 0
	}
	err = global.DB.Delete(&mc).Error
	if err != nil {
		return 0
	}
	return 1
}
