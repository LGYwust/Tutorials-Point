package models

type MovieModel struct {
	Model
	Name     string `gorm:"column:name" json:"name"`         //中文名
	Image    string `gorm:"column:image" json:"image"`       //图片_url
	Director string `gorm:"column:director" json:"director"` //导演
	Time     string `gorm:"column:time" json:"time"`         //时间
	Country  string `gorm:"column:country" json:"country"`   //国家
	Category string `gorm:"column:category" json:"category"` //分类
	Rating   string `gorm:"column:rating" json:"rating"`     //评分
	People   int    `gorm:"column:people" json:"people"`     //有多少人评价
	Quote    string `gorm:"column:quote" json:"quote"`       //对应的一句话
	Url      string `gorm:"column:url" json:"url"`           //电影地址
}
