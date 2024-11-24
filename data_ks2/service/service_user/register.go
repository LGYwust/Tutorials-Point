package service_user

import (
	"data_ks2/global"
	"data_ks2/models"
	"data_ks2/utils/pwd"
)

func CheckUserName(account string) bool {
	var user models.UserModel
	err := global.DB.Take(&user, "account =?", account).Error
	if err != nil {
		return false
	}
	return true
}

func AddUser(account string, email string, phone string, password string) bool {
	err := global.DB.Create(&models.UserModel{
		Account:  account,
		Password: pwd.HashPassword(password),
		Phone:    phone,
		Email:    email,
	}).Error
	if err != nil {
		return true
	}
	return false
}
