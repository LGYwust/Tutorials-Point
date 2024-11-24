package models

type UserModel struct {
	Model
	Account  string `gorm:"column:account;size:64" json:"account"`
	Email    string `gorm:"column:email;size:128" json:"email"`
	Phone    string `gorm:"column:phone;size:64" json:"phone"`
	Password string `gorm:"column:password;size:64" json:"password"`
	RoleID   int    `gorm:"column:roleid;size:64" json:"roleid"`
	NickName string `gorm:"column:nickname;size:64" json:"nickname"`
	Old      int    `gorm:"column:old;size:64" json:"old"`
}
