package service_user

import (
	"data_ks2/global"
	"data_ks2/models"
	"data_ks2/utils/jwts"
	"data_ks2/utils/pwd"
)

func LoginAccount(roleid int, account string, password string) (string, bool) {
	var user models.UserModel
	err := global.DB.Take(&user, "account =?", account).Error
	if err != nil {
		return "", false
	}
	if roleid == user.RoleID || (roleid == 1 && user.RoleID == 2) {
		if pwd.CheckPassword(user.Password, password) {
			token, err := jwts.GenToken(jwts.JwyPayLoad{
				NickName: user.NickName,
				RoleID:   user.RoleID,
				UserID:   user.ID,
			})
			if err != nil {
				return "", false
			}
			return token, true
		}
	}
	return "", false
}
