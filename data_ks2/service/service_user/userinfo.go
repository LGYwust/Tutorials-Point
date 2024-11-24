package service_user

import (
	"data_ks2/global"
	"data_ks2/models"
)

func SaveUserinfo(userinfo models.UserinfoModel) bool {
	var ui models.UserinfoModel
	err := global.DB.Take(&ui, "userid =?", userinfo.UserID).Error
	if err != nil {
		if err := global.DB.Create(&userinfo).Error; err != nil {
			return false
		}
	} else {
		// 如果找到记录，更新用户信息
		if err := global.DB.Model(&ui).Updates(models.UserinfoModel{
			NickName: userinfo.NickName,
			Time:     userinfo.Time,
			Sex:      userinfo.Sex,
			Desc:     userinfo.Desc,
		}).Error; err != nil {
			return false
		}
	}
	return true
}

func Getuserinfo(userid int) models.UserinfoModel {
	var ui models.UserinfoModel
	err := global.DB.Take(&ui, "userid = ?", userid).Error
	if err != nil {
		return models.UserinfoModel{}
	}
	return ui
}

func Getalluser() []models.UserModel {
	var um []models.UserModel
	err := global.DB.Find(&um).Error
	if err != nil {
		return nil
	}
	return um
}
