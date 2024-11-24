package models

type ConsultationModel struct {
	Model
	Image       string `gorm:"column:image" json:"image"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	Source      string `gorm:"column:source" json:"source"`
	ReleaseDate string `gorm:"column:releaseDate" json:"releaseDate"`
	Url         string `gorm:"column:url" json:"url"`
}
