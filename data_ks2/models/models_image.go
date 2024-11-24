package models

type ImageModel struct {
	Model
	UserID    int       `gorm:"column:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	FileName  string    `gorm:"column:fileName" json:"fileName"`
	Size      int64     `gorm:"column:size" json:"size"`
	Path      string    `gorm:"column:path" json:"path"`
	Hash      string    `gorm:"column:hash" json:"hash"`
}
