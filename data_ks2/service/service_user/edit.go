package service_user

import (
	"data_ks2/global"
	"data_ks2/models"
)

func EditUser(Account string, RoleID int, userID int) int {
	if userID == 2 {
		var um models.UserModel
		err := global.DB.Take(&um, "account = ?", Account).Error
		if err != nil {
			return 0
		}
		// 如果找到记录，更新用户信息
		um.RoleID = RoleID
		if err := global.DB.Save(&um).Error; err != nil {
			return 0
		}
		return 1
	}
	return -1
}

func DeltUser(account string, roleid int) int {
	var um models.UserModel
	err := global.DB.Take(&um, "account = ?", account).Error
	if err != nil {
		return 0
	}
	if (um.RoleID == 1 && roleid != 2) || um.RoleID == 2 {
		return -1
	}
	err = global.DB.Delete(&um).Error
	if err != nil {
		return 0
	}
	return 1
}
