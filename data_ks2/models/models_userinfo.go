package models

type UserinfoModel struct {
	Model
	UserID    int       `gorm:"column:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	NickName  string    `gorm:"column:nickname;size:64" json:"nickname"`
	Time      string    `gorm:"column:time;size:32" json:"time"`
	Sex       string    `gorm:"column:sex;size:32" json:"sex"`
	Desc      string    `gorm:"column:desc;size:256" json:"desc"`
}
